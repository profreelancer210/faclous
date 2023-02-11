// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/function/errors"
	"github.com/ysugimoto/falco/interpreter/value"
)

var Std_collect_ArgumentTypes = []value.Type{value.IdentType, value.StringType}

func Std_collect_Validate(args []value.Value) error {
	if len(args) < 1 || len(args) > 2 {
		return errors.ArgumentNotInRange("std.collect", 1, 2, args)
	}
	for i := range args {
		if args[i].Type() != Std_collect_ArgumentTypes[i] {
			return errors.TypeMismatch("std.collect", i+1, Std_collect_ArgumentTypes[i], args[i].Type())
		}
	}
	return nil
}

// Fastly built-in function implementation of std.collect
// Arguments may be:
// - ID
// - ID, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/miscellaneous/std-collect/
func Std_collect(ctx *context.Context, args ...value.Value) (value.Value, error) {
	// Argument validations
	if err := Std_collect_Validate(args); err != nil {
		return value.Null, err
	}

	// Need to be implemented
	return value.Null, errors.NotImplemented("std.collect")
}
