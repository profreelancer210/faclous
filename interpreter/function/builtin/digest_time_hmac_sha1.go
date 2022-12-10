// Code generated by __generator__/interpreter.go; DO NOT EDIT.

package builtin

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/value"
)

// Fastly built-in function implementation of digest.time_hmac_sha1
// Arguments may be:
// - STRING, INTEGER, INTEGER
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/digest-time-hmac-sha1/
func Digest_time_hmac_sha1(ctx *context.Context, args ...value.Value) (value.Value, error) {
	// Need to be implemented
	return value.Null, errors.WithStack(fmt.Errorf("Not Impelemented"))
}
