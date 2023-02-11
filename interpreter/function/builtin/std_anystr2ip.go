// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/function/errors"
	"github.com/ysugimoto/falco/interpreter/value"
)

var Std_anystr2ip_ArgumentTypes = []value.Type{value.StringType, value.StringType}

func Std_anystr2ip_Validate(args []value.Value) error {
	if len(args) != 2 {
		return errors.ArgumentNotEnough("std.anystr2ip", 2, args)
	}
	for i := range args {
		if args[i].Type() != Std_anystr2ip_ArgumentTypes[i] {
			return errors.TypeMismatch("std.anystr2ip", i+1, Std_anystr2ip_ArgumentTypes[i], args[i].Type())
		}
	}
	return nil
}

// Fastly built-in function implementation of std.anystr2ip
// Arguments may be:
// - STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-anystr2ip/
func Std_anystr2ip(ctx *context.Context, args ...value.Value) (value.Value, error) {
	// Argument validations
	if err := Std_anystr2ip_Validate(args); err != nil {
		return value.Null, err
	}

	// Need to be implemented
	return value.Null, errors.NotImplemented("std.anystr2ip")
}
