// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"time"

	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/function/errors"
	"github.com/ysugimoto/falco/interpreter/value"
)

const Std_integer2time_Name = "std.integer2time"

var Std_integer2time_ArgumentTypes = []value.Type{value.IntegerType}

func Std_integer2time_Validate(args []value.Value) error {
	if len(args) != 1 {
		return errors.ArgumentNotEnough(Std_integer2time_Name, 1, args)
	}
	for i := range args {
		if args[i].Type() != Std_integer2time_ArgumentTypes[i] {
			return errors.TypeMismatch(Std_integer2time_Name, i+1, Std_integer2time_ArgumentTypes[i], args[i].Type())
		}
	}
	return nil
}

// Fastly built-in function implementation of std.integer2time
// Arguments may be:
// - INTEGER
// Reference: https://developer.fastly.com/reference/vcl/functions/date-and-time/std-integer2time/
func Std_integer2time(ctx *context.Context, args ...value.Value) (value.Value, error) {
	// Argument validations
	if err := Std_integer2time_Validate(args); err != nil {
		return value.Null, err
	}

	t := value.Unwrap[*value.Integer](args[0])

	return &value.Time{Value: time.Unix(t.Value, 0).UTC()}, nil
}
