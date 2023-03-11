// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"math/rand"
	"time"

	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/function/errors"
	"github.com/ysugimoto/falco/interpreter/value"
)

const Randomstr_Name = "randomstr"

var Randomstr_ArgumentTypes = []value.Type{value.IntegerType, value.StringType}

func Randomstr_Validate(args []value.Value) error {
	if len(args) < 1 || len(args) > 2 {
		return errors.ArgumentNotInRange(Randomstr_Name, 1, 2, args)
	}
	for i := range args {
		if args[i].Type() != Randomstr_ArgumentTypes[i] {
			return errors.TypeMismatch(Randomstr_Name, i+1, Randomstr_ArgumentTypes[i], args[i].Type())
		}
	}
	return nil
}

// Fastly built-in function implementation of randomstr
// Arguments may be:
// - INTEGER
// - INTEGER, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/randomness/randomstr/
func Randomstr(ctx *context.Context, args ...value.Value) (value.Value, error) {
	// Argument validations
	if err := Randomstr_Validate(args); err != nil {
		return value.Null, err
	}

	length := value.Unwrap[*value.Integer](args[0])
	characters := []rune(value.Unwrap[*value.String](args[1]).Value)

	rand.Seed(time.Now().UnixNano())
	ret := make([]rune, int(length.Value))

	for i := 0; i < int(length.Value); i++ {
		ret[i] = characters[rand.Intn(len(characters)-1)]
	}

	return &value.String{Value: string(ret)}, nil
}
