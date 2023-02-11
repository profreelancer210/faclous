// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/function/errors"
	"github.com/ysugimoto/falco/interpreter/value"
)

var Digest_base64url_ArgumentTypes = []value.Type{value.StringType}

func Digest_base64url_Validate(args []value.Value) error {
	if len(args) != 1 {
		return errors.ArgumentNotEnough("digest.base64url", 1, args)
	}
	for i := range args {
		if args[i].Type() != Digest_base64url_ArgumentTypes[i] {
			return errors.TypeMismatch("digest.base64url", i+1, Digest_base64url_ArgumentTypes[i], args[i].Type())
		}
	}
	return nil
}

// Fastly built-in function implementation of digest.base64url
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/digest-base64url/
func Digest_base64url(ctx *context.Context, args ...value.Value) (value.Value, error) {
	// Argument validations
	if err := Digest_base64url_Validate(args); err != nil {
		return value.Null, err
	}

	// Need to be implemented
	return value.Null, errors.NotImplemented("digest.base64url")
}
