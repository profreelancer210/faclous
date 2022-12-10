// Code generated by __generator__/interpreter.go; DO NOT EDIT.

package builtin

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/value"
)

// Fastly built-in function implementation of crypto.encrypt_base64
// Arguments may be:
// - ID, ID, ID, STRING, STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/crypto-encrypt-base64/
func Crypto_encrypt_base64(ctx *context.Context, args ...value.Value) (value.Value, error) {
	// Need to be implemented
	return value.Null, errors.WithStack(fmt.Errorf("Not Impelemented"))
}
