// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"

	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/function/errors"
	"github.com/ysugimoto/falco/interpreter/value"
)

const Digest_hash_sha1_from_base64_Name = "digest.hash_sha1_from_base64"

var Digest_hash_sha1_from_base64_ArgumentTypes = []value.Type{value.StringType}

func Digest_hash_sha1_from_base64_Validate(args []value.Value) error {
	if len(args) != 1 {
		return errors.ArgumentNotEnough(Digest_hash_sha1_from_base64_Name, 1, args)
	}
	for i := range args {
		if args[i].Type() != Digest_hash_sha1_from_base64_ArgumentTypes[i] {
			return errors.TypeMismatch(Digest_hash_sha1_from_base64_Name, i+1, Digest_hash_sha1_from_base64_ArgumentTypes[i], args[i].Type())
		}
	}
	return nil
}

// Fastly built-in function implementation of digest.hash_sha1_from_base64
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/digest-hash-sha1-from-base64/
func Digest_hash_sha1_from_base64(ctx *context.Context, args ...value.Value) (value.Value, error) {
	// Argument validations
	if err := Digest_hash_sha1_from_base64_Validate(args); err != nil {
		return value.Null, err
	}

	input := value.Unwrap[*value.String](args[0])
	decoded, err := base64.StdEncoding.DecodeString(input.Value)
	if err != nil {
		return value.Null, err
	}
	enc := sha1.Sum(decoded)

	return &value.String{
		Value: hex.EncodeToString(enc[:]),
	}, nil
}
