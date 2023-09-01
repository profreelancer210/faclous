package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strings"

	"encoding/json"

	"github.com/fatih/color"
	"github.com/kyokomi/emoji"
	"github.com/mattn/go-colorable"
	"github.com/pkg/errors"
	"github.com/ysugimoto/falco/config"
	ife "github.com/ysugimoto/falco/interpreter/function/errors"
	"github.com/ysugimoto/falco/lexer"
	"github.com/ysugimoto/falco/resolver"
	"github.com/ysugimoto/falco/terraform"
	"github.com/ysugimoto/falco/tester"
	"github.com/ysugimoto/falco/token"
)

var version string = ""

var (
	output  = colorable.NewColorableStderr()
	yellow  = color.New(color.FgYellow)
	white   = color.New(color.FgWhite)
	red     = color.New(color.FgRed)
	green   = color.New(color.FgGreen)
	cyan    = color.New(color.FgCyan)
	magenta = color.New(color.FgMagenta)

	// Displaying test result needs adding background colors
	noTestColor = color.New(color.FgBlack, color.BgWhite, color.Bold)
	passColor   = color.New(color.FgWhite, color.BgGreen, color.Bold)
	failColor   = color.New(color.FgWhite, color.BgRed, color.Bold)
	redBold     = color.New(color.FgRed, color.Bold)

	ErrExit = errors.New("exit")
)

const (
	subcommandLint      = "lint"
	subcommandTerraform = "terraform"
	subcommandSimulate  = "simulate"
	subcommandStats     = "stats"
	subcommandTest      = "test"
)

func write(c *color.Color, format string, args ...interface{}) {
	c.Fprint(output, emoji.Sprintf(format, args...))
}

func writeln(c *color.Color, format string, args ...interface{}) {
	write(c, format+"\n", args...)
}

func main() {
	c, err := config.New(os.Args[1:])
	if err != nil {
		writeln(red, "Failed to initialize config: %s", err)
		os.Exit(1)
	}
	if c.Help {
		printHelp(c.Commands.At(0))
		os.Exit(1)
	} else if c.Version {
		writeln(white, version)
		os.Exit(1)
	}

	var fetcher Fetcher
	var action string
	// falco could lint multiple services so resolver should be a slice
	var resolvers []resolver.Resolver
	switch c.Commands.At(0) {
	case subcommandTerraform:
		fastlyServices, err := ParseStdin()
		if err == nil {
			resolvers = resolver.NewTerraformResolver(fastlyServices)
			fetcher = terraform.NewTerraformFetcher(fastlyServices)
		}
		action = c.Commands.At(1)
	case subcommandSimulate, subcommandLint, subcommandStats, subcommandTest:
		// "lint", "simulate", "stats" and "test" command provides single file of service,
		// then resolvers size is always 1
		resolvers, err = resolver.NewFileResolvers(c.Commands.At(1), c.IncludePaths)
		action = c.Commands.At(0)
	default:
		if filepath.Ext(c.Commands.At(0)) != ".vcl" {
			err = fmt.Errorf("Unrecognized subcommand: %s", c.Commands.At(0))
		} else {
			// "lint" command provides single file of service, then resolvers size is always 1
			resolvers, err = resolver.NewFileResolvers(c.Commands.At(0), c.IncludePaths)
			action = c.Commands.At(0)
		}
	}

	if err != nil {
		writeln(red, err.Error())
		os.Exit(1)
	}

	var shouldExit bool
	for _, v := range resolvers {
		runner, err := NewRunner(c, fetcher)
		if err != nil {
			writeln(red, err.Error())
			os.Exit(1)
		}

		var exitErr error
		switch action {
		case subcommandTest:
			exitErr = runTest(runner, v)
		case subcommandSimulate:
			exitErr = runSimulate(runner, v)
		case subcommandStats:
			exitErr = runStats(runner, v)
		default:
			if name := v.Name(); name != "" {
				writeln(white, `Lint service of "%s"`, name)
				writeln(white, strings.Repeat("=", 18+len(name)))
			}
			exitErr = runLint(runner, v)
		}

		if exitErr == ErrExit {
			shouldExit = true
		}
	}

	if shouldExit {
		os.Exit(1)
	}
}

func runLint(runner *Runner, rslv resolver.Resolver) error {
	result, err := runner.Run(rslv)
	if err != nil {
		if err != ErrParser {
			writeln(red, err.Error())
		}
		return ErrExit
	}

	if runner.config.Json {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		if err := enc.Encode(result); err != nil {
			writeln(red, err.Error())
			return ErrExit
		}
	}

	write(red, ":fire:%d errors, ", result.Errors)
	write(yellow, ":exclamation:%d warnings, ", result.Warnings)
	writeln(cyan, ":speaker:%d recommendations.", result.Infos)

	// Display message corresponds to runner result
	if result.Errors == 0 {
		switch {
		case result.Warnings > 0:
			writeln(white, "VCL lint warnings encountered, but things should run OK :thumbsup:")
			if runner.level < LevelWarning {
				writeln(white, "Run command with the -v option to output warnings.")
			}
		case result.Infos > 0:
			writeln(green, "VCL looks good :sparkles: Some recommendations are available :thumbsup:")
			if runner.level < LevelInfo {
				writeln(white, "Run command with the -vv option to output recommendations.")
			}
		default:
			writeln(green, "VCL looks great :sparkles:")
		}
	}

	// if lint error is not zero, stop process
	if result.Errors > 0 {
		if len(runner.transformers) > 0 {
			writeln(white, "Program aborted. Please fix lint errors before transforming.")
		}
		return ErrExit
	}

	if err := runner.Transform(result.Vcl); err != nil {
		writeln(red, err.Error())
		return ErrExit
	}
	return nil
}

func runSimulate(runner *Runner, rslv resolver.Resolver) error {
	if err := runner.Simulate(rslv); err != nil {
		writeln(red, "Failed to start local simulator: %s", err.Error())
		return ErrExit
	}
	return nil
}

func runStats(runner *Runner, rslv resolver.Resolver) error {
	stats, err := runner.Stats(rslv)
	if err != nil {
		if err != ErrParser {
			writeln(red, err.Error())
		}
		return ErrExit
	}

	if runner.config.Json {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		if err := enc.Encode(stats); err != nil {
			writeln(red, err.Error())
			return ErrExit
		}
		return nil
	}
	printStats := func(format string, args ...interface{}) {
		fmt.Fprintf(os.Stdout, format+"\n", args...)
	}

	printStats(strings.Repeat("=", 80))
	printStats("| %-76s |", "falco VCL statistics ")
	printStats(strings.Repeat("=", 80))
	printStats("| %-22s | %51s |", "Main VCL File", stats.Main)
	printStats(strings.Repeat("=", 80))
	printStats("| %-22s | %51d |", "Included Module Files", stats.Files-1)
	printStats(strings.Repeat("-", 80))
	printStats("| %-22s | %51d |", "Total VCL Lines", stats.Lines)
	printStats(strings.Repeat("-", 80))
	printStats("| %-22s | %51d |", "Subroutines", stats.Subroutines)
	printStats(strings.Repeat("-", 80))
	printStats("| %-22s | %51d |", "Backends", stats.Backends)
	printStats(strings.Repeat("-", 80))
	printStats("| %-22s | %51d |", "Tables", stats.Tables)
	printStats(strings.Repeat("-", 80))
	printStats("| %-22s | %51d |", "Access Control Lists", stats.Acls)
	printStats(strings.Repeat("-", 80))
	printStats("| %-22s | %51d |", "Directors", stats.Directors)
	printStats(strings.Repeat("-", 80))
	return nil
}

func runTest(runner *Runner, rslv resolver.Resolver) error {
	factory, err := runner.Test(rslv)
	if err != nil {
		return ErrExit
	}

	if runner.config.Json {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		if err := enc.Encode(struct {
			Tests   []*tester.TestResult `json:"tests"`
			Summary *tester.TestCounter  `json:"summary"`
		}{
			Tests:   factory.Results,
			Summary: factory.Statistics,
		}); err != nil {
			writeln(red, err.Error())
			return ErrExit
		}
		return nil
	}

	// shrthand indent making
	indent := func(level int) string {
		return strings.Repeat(" ", level*2)
	}
	// print problem line
	printCodeLine := func(lx *lexer.Lexer, tok token.Token) {
		problemLine := tok.Line
		lineFormat := fmt.Sprintf(" %%%dd", int(math.Floor(math.Log10(float64(problemLine+1))+1)))
		for l := problemLine - 1; l <= problemLine+1; l++ {
			line, ok := lx.GetLine(l)
			if !ok {
				continue
			}
			color := white
			if l == problemLine {
				color = yellow
			}
			writeln(color, "%s "+lineFormat+"| %s", indent(1), l, strings.ReplaceAll(line, "\t", "    "))
		}
	}

	var passedCount, failedCount, totalCount int
	for _, r := range factory.Results {
		switch {
		case len(r.Cases) == 0:
			write(noTestColor, " NO TESTS ")
			writeln(white, " "+r.Filename)
		case r.IsPassed():
			write(passColor, " PASS ")
			writeln(white, " "+r.Filename)
		default:
			write(failColor, " FAIL ")
			writeln(white, " "+r.Filename)
		}

		for _, c := range r.Cases {
			totalCount++
			if c.Error != nil {
				writeln(redBold, "%s●  [%s] %s\n", indent(1), c.Scope, c.Name)
				writeln(red, "%s%s", indent(2), c.Error.Error())
				switch e := c.Error.(type) {
				case *ife.AssertionError:
					write(white, "%sActual Value: ", indent(2))
					writeln(red, "%s\n", e.Actual.String())
					printCodeLine(r.Lexer, e.Token)
				case *ife.TestingError:
					writeln(white, "")
					printCodeLine(r.Lexer, e.Token)
				}
				writeln(white, "")
				failedCount++
			} else {
				writeln(green, "%s✓ [%s] %s", indent(1), c.Scope, c.Name)
				passedCount++
			}
		}
	}

	if passedCount > 0 {
		write(green, "%d passed, ", passedCount)
	} else {
		write(white, "%d passed, ", passedCount)
	}
	if failedCount > 0 {
		write(red, "%d failed, ", failedCount)
	} else {
		write(white, "%d failed, ", failedCount)
	}
	write(white, "%d total, ", totalCount)
	writeln(white, "%d assertions", factory.Statistics.Asserts)

	return nil
}
