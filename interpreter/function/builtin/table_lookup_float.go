// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"github.com/ysugimoto/falco/ast"
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/function/errors"
	"github.com/ysugimoto/falco/interpreter/value"
)

const Table_lookup_float_Name = "table.lookup_float"

var Table_lookup_float_ArgumentTypes = []value.Type{value.IdentType, value.StringType, value.FloatType}

func Table_lookup_float_Validate(args []value.Value) error {
	if len(args) != 3 {
		return errors.ArgumentNotEnough(Table_lookup_float_Name, 3, args)
	}
	for i := range args {
		if args[i].Type() != Table_lookup_float_ArgumentTypes[i] {
			return errors.TypeMismatch(Table_lookup_float_Name, i+1, Table_lookup_float_ArgumentTypes[i], args[i].Type())
		}
	}
	return nil
}

// Fastly built-in function implementation of table.lookup_float
// Arguments may be:
// - TABLE, STRING, FLOAT
// Reference: https://developer.fastly.com/reference/vcl/functions/table/table-lookup-float/
func Table_lookup_float(ctx *context.Context, args ...value.Value) (value.Value, error) {
	// Argument validations
	if err := Table_lookup_float_Validate(args); err != nil {
		return value.Null, err
	}

	id := value.Unwrap[*value.Ident](args[0]).Value
	key := value.Unwrap[*value.String](args[1]).Value
	defaultValue := value.Unwrap[*value.Float](args[2]).Value

	table, ok := ctx.Tables[id]
	if !ok {
		return &value.Float{Value: defaultValue}, errors.New(Table_lookup_float_Name,
			"table %d does not exist", id,
		)
	}
	if table.ValueType == nil || table.ValueType.Value != "FLOAT" {
		return &value.Float{Value: defaultValue}, errors.New(Table_lookup_float_Name,
			"table %d value type is not FLOAT", id,
		)
	}

	for _, prop := range table.Properties {
		if prop.Key.Value == key {
			v, ok := prop.Value.(*ast.Float)
			if !ok {
				return &value.Float{Value: defaultValue}, errors.New(Table_lookup_float_Name,
					"table %s value could not cast to FLOAT type", id,
				)
			}
			return &value.Float{Value: v.Value}, nil
		}
	}
	return &value.Float{Value: defaultValue}, nil
}
