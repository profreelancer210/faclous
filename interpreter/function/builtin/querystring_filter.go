// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/function/errors"
	"github.com/ysugimoto/falco/interpreter/value"
)

const Querystring_filter_Name = "querystring.filter"

var Querystring_filter_ArgumentTypes = []value.Type{value.StringType, value.StringType}

func Querystring_filter_Validate(args []value.Value) error {
	if len(args) != 2 {
		return errors.ArgumentNotEnough(Querystring_filter_Name, 2, args)
	}
	for i := range args {
		if args[i].Type() != Querystring_filter_ArgumentTypes[i] {
			return errors.TypeMismatch(Querystring_filter_Name, i+1, Querystring_filter_ArgumentTypes[i], args[i].Type())
		}
	}
	return nil
}

// Fastly built-in function implementation of querystring.filter
// Arguments may be:
// - STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/query-string/querystring-filter/
func Querystring_filter(ctx *context.Context, args ...value.Value) (value.Value, error) {
	// Argument validations
	if err := Querystring_filter_Validate(args); err != nil {
		return value.Null, err
	}

	// Need to be implemented
	return value.Null, errors.NotImplemented("querystring.filter")
}
