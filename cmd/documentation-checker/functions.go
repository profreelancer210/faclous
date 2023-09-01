package main

import (
	"context"
	"os"
	"sync"

	"github.com/go-yaml/yaml"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

var fastlyFunctionCategoryPageUrls = []string{
	"/reference/vcl/functions/content-negotiation/",
	"/reference/vcl/functions/cryptographic/",
	"/reference/vcl/functions/date-and-time/",
	"/reference/vcl/functions/floating-point-classifications/",
	"/reference/vcl/functions/headers/",
	"/reference/vcl/functions/math-logexp/",
	"/reference/vcl/functions/math-rounding/",
	"/reference/vcl/functions/math-trig/",
	"/reference/vcl/functions/miscellaneous/",
	"/reference/vcl/functions/query-string/",
	"/reference/vcl/functions/randomness/",
	"/reference/vcl/functions/rate-limiting/",
	"/reference/vcl/functions/strings/",
	"/reference/vcl/functions/tls-and-http/",
	"/reference/vcl/functions/table/",
	"/reference/vcl/functions/uuid/",
}

const builtinPath = "../../__generator__/builtin.yml"

func factoryFunctions(ctx context.Context) (*sync.Map, error) {
	var eg errgroup.Group
	var m sync.Map
	for i := range fastlyFunctionCategoryPageUrls {
		url := fastlyDocDomain + fastlyFunctionCategoryPageUrls[i]
		eg.Go(func() error {
			return fetchFastlyDocument(ctx, url, &m)
		})
	}
	if err := eg.Wait(); err != nil {
		return &m, errors.WithStack(err)
	}
	return &m, nil
}

func checkFunctions(m *sync.Map) error {
	fp, err := os.Open(builtinPath)
	if err != nil {
		return errors.WithStack(err)
	}
	defer fp.Close()

	variables := make(map[string]interface{})
	if err := yaml.NewDecoder(fp).Decode(variables); err != nil {
		return errors.WithStack(err)
	}

	m.Range(func(key, val interface{}) bool {
		k := key.(string) // nolint:errcheck
		v := val.(string) // nolint:errcheck

		if _, ok := variables[k]; ok {
			return true
		}
		write(yellow, "[!] ")
		writeln(white, `"%s" is not defined, url: %s`, k, v)
		return true
	})
	return nil
}
