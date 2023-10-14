// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"github.com/google/uuid"
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/function/errors"
	"github.com/ysugimoto/falco/interpreter/value"
)

const Uuid_oid_Name = "uuid.oid"

var Uuid_oid_ArgumentTypes = []value.Type{}

func Uuid_oid_Validate(args []value.Value) error {
	if len(args) > 0 {
		return errors.ArgumentMustEmpty(Uuid_oid_Name, args)
	}
	return nil
}

// Fastly built-in function implementation of uuid.oid
// Arguments may be:
// Reference: https://developer.fastly.com/reference/vcl/functions/uuid/uuid-oid/
func Uuid_oid(ctx *context.Context, args ...value.Value) (value.Value, error) {
	// Argument validations
	if err := Uuid_oid_Validate(args); err != nil {
		return value.Null, err
	}

	// ISO OID namespace, namely constant string
	return &value.String{Value: uuid.NameSpaceOID.String()}, nil
}
