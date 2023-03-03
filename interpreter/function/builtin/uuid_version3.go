// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"github.com/google/uuid"
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/function/errors"
	"github.com/ysugimoto/falco/interpreter/value"
)

const Uuid_version3_Name = "uuid.version3"

var Uuid_version3_ArgumentTypes = []value.Type{value.StringType, value.StringType}

func Uuid_version3_Validate(args []value.Value) error {
	if len(args) != 2 {
		return errors.ArgumentNotEnough(Uuid_version3_Name, 2, args)
	}
	for i := range args {
		if args[i].Type() != Uuid_version3_ArgumentTypes[i] {
			return errors.TypeMismatch(Uuid_version3_Name, i+1, Uuid_version3_ArgumentTypes[i], args[i].Type())
		}
	}
	return nil
}

// Fastly built-in function implementation of uuid.version3
// Arguments may be:
// - STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/uuid/uuid-version3/
func Uuid_version3(ctx *context.Context, args ...value.Value) (value.Value, error) {
	// Argument validations
	if err := Uuid_version3_Validate(args); err != nil {
		return value.Null, err
	}

	namespace := value.Unwrap[*value.String](args[0]).Value
	name := value.Unwrap[*value.String](args[1]).Value

	space, err := uuid.Parse(namespace)
	if err != nil {
		return &value.String{IsNotSet: true}, errors.New(Uuid_version3_Name,
			"Failed to parse namespace of %s", namespace,
		)
	}

	return &value.String{
		Value: uuid.NewMD5(space, []byte(name)).String(),
	}, nil
}
