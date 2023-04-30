// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"github.com/google/uuid"
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/function/errors"
	"github.com/ysugimoto/falco/interpreter/value"
)

const Uuid_is_version5_Name = "uuid.is_version5"

var Uuid_is_version5_ArgumentTypes = []value.Type{value.StringType}

func Uuid_is_version5_Validate(args []value.Value) error {
	if len(args) != 1 {
		return errors.ArgumentNotEnough(Uuid_is_version5_Name, 1, args)
	}
	for i := range args {
		if args[i].Type() != Uuid_is_version5_ArgumentTypes[i] {
			return errors.TypeMismatch(Uuid_is_version5_Name, i+1, Uuid_is_version5_ArgumentTypes[i], args[i].Type())
		}
	}
	return nil
}

// Fastly built-in function implementation of uuid.is_version5
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/uuid/uuid-is-version5/
func Uuid_is_version5(ctx *context.Context, args ...value.Value) (value.Value, error) {
	// Argument validations
	if err := Uuid_is_version5_Validate(args); err != nil {
		return value.Null, err
	}

	input := value.Unwrap[*value.String](args[0]).Value
	if id, err := uuid.Parse(input); err != nil {
		return &value.Boolean{Value: false}, nil
	} else if id.Version() != 5 {
		return &value.Boolean{Value: false}, nil
	}
	return &value.Boolean{Value: true}, nil
}
