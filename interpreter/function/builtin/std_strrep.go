// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/function/errors"
	"github.com/ysugimoto/falco/interpreter/value"
)

const Std_strrep_Name = "std.strrep"

var Std_strrep_ArgumentTypes = []value.Type{value.StringType, value.IntegerType}

func Std_strrep_Validate(args []value.Value) error {
	if len(args) != 2 {
		return errors.ArgumentNotEnough(Std_strrep_Name, 2, args)
	}
	for i := range args {
		if args[i].Type() != Std_strrep_ArgumentTypes[i] {
			return errors.TypeMismatch(Std_strrep_Name, i+1, Std_strrep_ArgumentTypes[i], args[i].Type())
		}
	}
	return nil
}

// Fastly built-in function implementation of std.strrep
// Arguments may be:
// - STRING, INTEGER
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-strrep/
func Std_strrep(ctx *context.Context, args ...value.Value) (value.Value, error) {
	// Argument validations
	if err := Std_strrep_Validate(args); err != nil {
		return value.Null, err
	}

	// Need to be implemented
	return value.Null, errors.NotImplemented("std.strrep")
}
