// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/function/errors"
	"github.com/ysugimoto/falco/interpreter/value"
)

const Table_lookup_bool_Name = "table.lookup_bool"

var Table_lookup_bool_ArgumentTypes = []value.Type{value.IdentType, value.StringType}

func Table_lookup_bool_Validate(args []value.Value) error {
	if len(args) < 2 || len(args) > 3 {
		return errors.ArgumentNotInRange(Table_lookup_bool_Name, 2, 3, args)
	}
	for i := range args {
		if args[i].Type() != Table_lookup_bool_ArgumentTypes[i] {
			return errors.TypeMismatch(Table_lookup_bool_Name, i+1, Table_lookup_bool_ArgumentTypes[i], args[i].Type())
		}
	}
	return nil
}

// Fastly built-in function implementation of table.lookup_bool
// Arguments may be:
// - TABLE, STRING, BOOL
// - TABLE, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/table/table-lookup-bool/
func Table_lookup_bool(ctx *context.Context, args ...value.Value) (value.Value, error) {
	// Argument validations
	if err := Table_lookup_bool_Validate(args); err != nil {
		return value.Null, err
	}

	// Need to be implemented
	return value.Null, errors.NotImplemented("table.lookup_bool")
}
