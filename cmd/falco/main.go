package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"encoding/json"

	"github.com/fatih/color"
	"github.com/kyokomi/emoji"
	"github.com/mattn/go-colorable"
	"github.com/pkg/errors"
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

	ErrExit = errors.New("exit")
)

const (
	subcommandLint      = "lint"
	subcommandTerraform = "terraform"
)

type multiStringFlags []string

func (m *multiStringFlags) Set(v string) error {
	*m = append(*m, v)
	return nil
}

func (m *multiStringFlags) String() string {
	return fmt.Sprintf("%v", *m)
}

type Config struct {
	Transforms   multiStringFlags
	IncludePaths multiStringFlags
	Help         bool
	V            bool
	VV           bool
	Version      bool
	Remote       bool
	Stats        bool
	Json         bool
}

func write(c *color.Color, format string, args ...interface{}) {
	c.Fprint(output, emoji.Sprintf(format, args...))
}

func writeln(c *color.Color, format string, args ...interface{}) {
	write(c, format+"\n", args...)
}

func printUsage() {
	usage := `
=======================================
  falco: Fastly VCL parser / linter
=======================================
Usage:
    falco [subcommand] [main vcl file]

Subcommands:
    terraform : Run lint from terraform planned JSON
    lint      : Run lint (default)

Flags:
    -I, --include_path : Add include path
    -t, --transformer  : Specify transformer
    -h, --help         : Show this help
    -r, --remote       : Communicate with Fastly API
    -V, --version      : Display build version
    -v                 : Verbose warning lint result
    -vv                : Varbose all lint result
    -json              : Output statistics as JSON
    -stats             : Analyze VCL statistics

Simple Linting example:
    falco -I . -vv /path/to/vcl/main.vcl

Get statistics example:
    falco -I . -stats /path/to/vcl/main.vcl

Linting with terraform:
    terraform plan -out planned.out
    terraform show -json planned.out | falco -vv terraform
`

	fmt.Println(strings.TrimLeft(usage, "\n"))
	os.Exit(1)
}

func main() {
	c := &Config{}
	fs := flag.NewFlagSet("app", flag.ExitOnError)
	fs.Var(&c.IncludePaths, "I", "Add include paths (short)")
	fs.Var(&c.IncludePaths, "include_path", "Add include paths (long)")
	fs.Var(&c.Transforms, "t", "Add VCL transformer (short)")
	fs.Var(&c.Transforms, "transformer", "Add VCL transformer (long)")
	fs.BoolVar(&c.Help, "h", false, "Show Usage")
	fs.BoolVar(&c.Help, "help", false, "Show Usage")
	fs.BoolVar(&c.V, "v", false, "Verbose warning")
	fs.BoolVar(&c.VV, "vv", false, "Verbose info")
	fs.BoolVar(&c.Version, "V", false, "Print Version")
	fs.BoolVar(&c.Version, "version", false, "Print Version")
	fs.BoolVar(&c.Remote, "r", false, "Use Remote")
	fs.BoolVar(&c.Remote, "remote", false, "Use Remote")
	fs.BoolVar(&c.Stats, "stats", false, "Analyze VCL statistics")
	fs.BoolVar(&c.Json, "json", false, "Output statistics as JSON")
	fs.Usage = printUsage

	var err error
	if err = fs.Parse(os.Args[1:]); err != nil {
		if err == flag.ErrHelp {
			printUsage()
		}
		os.Exit(1)
	}

	if c.Help {
		printUsage()
	} else if c.Version {
		writeln(white, version)
		os.Exit(1)
	}

	// falco could lint multiple services so resolver should be a slice
	var resolvers []Resolver
	switch fs.Arg(0) {
	case subcommandTerraform:
		resolvers, err = NewStdinResolvers()
	case subcommandLint:
		// "lint" command provides single file of service, then resolvers size is always 1
		resolvers, err = NewFileResolvers(fs.Arg(1), c)
	default:
		// "lint" command provides single file of service, then resolvers size is always 1
		resolvers, err = NewFileResolvers(fs.Arg(0), c)
	}

	if err != nil {
		writeln(red, err.Error())
		os.Exit(1)
	}

	var shouldExit bool
	for _, v := range resolvers {
		if name := v.Name(); name != "" {
			writeln(white, `Lint service of "%s"`, name)
			writeln(white, strings.Repeat("=", 18+len(name)))
		}

		runner, err := NewRunner(v, c)
		if err != nil {
			writeln(red, err.Error())
			os.Exit(1)
		}

		var exitErr error
		if c.Stats {
			exitErr = runStats(runner, c.Json)
		} else {
			exitErr = runLint(runner)
		}
		if exitErr == ErrExit {
			shouldExit = true
		}
	}

	if shouldExit {
		os.Exit(1)
	}
}

func runLint(runner *Runner) error {
	result, err := runner.Run()
	if err != nil {
		if err != ErrParser {
			writeln(red, err.Error())
		}
		return ErrExit
	}

	write(red, ":fire:%d errors, ", result.Errors)
	write(yellow, ":exclamation:%d warnings, ", result.Warnings)
	writeln(cyan, ":speaker:%d infos.", result.Infos)

	// Display message corresponds to runner result
	if result.Errors == 0 {
		switch {
		case result.Warnings > 0:
			writeln(white, "VCL seems having some warnings, but it should be OK :thumbsup:")
			if runner.level < LevelWarning {
				writeln(white, "To see warning detail, run command with -v option.")
			}
		case result.Infos > 0:
			writeln(green, "VCL looks fine :sparkles: And we suggested some informations to vcl get more accuracy :thumbsup:")
			if runner.level < LevelInfo {
				writeln(white, "To see informations detail, run command with -vv option.")
			}
		default:
			writeln(green, "VCL looks very nice :sparkles:")
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

func runStats(runner *Runner, printJson bool) error {
	stats, err := runner.Stats()
	if err != nil {
		if err != ErrParser {
			writeln(red, err.Error())
		}
		return ErrExit
	}

	if printJson {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		if err := enc.Encode(stats); err != nil {
			writeln(red, err.Error())
			os.Exit(1)
		}
		return ErrExit
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

func printStats(format string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, format+"\n", args...)
}
