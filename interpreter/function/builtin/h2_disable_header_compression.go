// Code generated by __generator__/interpreter.go; DO NOT EDIT.

package builtin

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/value"
)

// Fastly built-in function implementation of h2.disable_header_compression
// Arguments may be:
// - STRING_LIST
// Reference: https://developer.fastly.com/reference/vcl/functions/tls-and-http/h2-disable-header-compression/
func H2_disable_header_compression(ctx *context.Context, args ...value.Value) (value.Value, error) {
	// Need to be implemented
	return value.Null, errors.WithStack(fmt.Errorf("Builtin function h2.disable_header_compression is not impelemented"))
}
