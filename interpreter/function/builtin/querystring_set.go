// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/function/errors"
	"github.com/ysugimoto/falco/interpreter/function/shared"
	"github.com/ysugimoto/falco/interpreter/value"
)

const Querystring_set_Name = "querystring.set"

var Querystring_set_ArgumentTypes = []value.Type{value.StringType, value.StringType, value.StringType}

func Querystring_set_Validate(args []value.Value) error {
	if len(args) != 3 {
		return errors.ArgumentNotEnough(Querystring_set_Name, 3, args)
	}
	for i := range args {
		if args[i].Type() != Querystring_set_ArgumentTypes[i] {
			return errors.TypeMismatch(Querystring_set_Name, i+1, Querystring_set_ArgumentTypes[i], args[i].Type())
		}
	}
	return nil
}

// Fastly built-in function implementation of querystring.set
// Arguments may be:
// - STRING, STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/query-string/querystring-set/
func Querystring_set(ctx *context.Context, args ...value.Value) (value.Value, error) {
	// Argument validations
	if err := Querystring_set_Validate(args); err != nil {
		return value.Null, err
	}

	u := value.Unwrap[*value.String](args[0])
	name := value.Unwrap[*value.String](args[1])
	val := value.Unwrap[*value.String](args[2])

	query, err := shared.ParseQuery(u.Value)
	if err != nil {
		return value.Null, errors.New(
			Querystring_set_Name, "Failed to parse urquery: %s, error: %s", u.Value, err.Error(),
		)
	}

	query.Set(name.Value, val.Value)
	return &value.String{Value: query.String()}, nil
}