// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/function/errors"
	"github.com/ysugimoto/falco/interpreter/value"
)

var Std_toupper_ArgumentTypes = []value.Type{value.StringType}

func Std_toupper_Validate(args []value.Value) error {
	if len(args) != 1 {
		return errors.ArgumentNotEnough("std.toupper", 1, args)
	}
	for i := range args {
		if args[i].Type() != Std_toupper_ArgumentTypes[i] {
			return errors.TypeMismatch("std.toupper", i+1, Std_toupper_ArgumentTypes[i], args[i].Type())
		}
	}
	return nil
}

// Fastly built-in function implementation of std.toupper
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-toupper/
func Std_toupper(ctx *context.Context, args ...value.Value) (value.Value, error) {
	// Argument validations
	if err := Std_toupper_Validate(args); err != nil {
		return value.Null, err
	}

	// Need to be implemented
	return value.Null, errors.NotImplemented("std.toupper")
}
