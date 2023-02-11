// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/function/errors"
	"github.com/ysugimoto/falco/interpreter/value"
)

var Time_add_ArgumentTypes = []value.Type{value.TimeType, value.RTimeType}

func Time_add_Validate(args []value.Value) error {
	if len(args) != 2 {
		return errors.ArgumentNotEnough("time.add", 2, args)
	}
	for i := range args {
		if args[i].Type() != Time_add_ArgumentTypes[i] {
			return errors.TypeMismatch("time.add", i+1, Time_add_ArgumentTypes[i], args[i].Type())
		}
	}
	return nil
}

// Fastly built-in function implementation of time.add
// Arguments may be:
// - TIME, RTIME
// Reference: https://developer.fastly.com/reference/vcl/functions/date-and-time/time-add/
func Time_add(ctx *context.Context, args ...value.Value) (value.Value, error) {
	// Argument validations
	if err := Time_add_Validate(args); err != nil {
		return value.Null, err
	}

	// Need to be implemented
	return value.Null, errors.NotImplemented("time.add")
}
