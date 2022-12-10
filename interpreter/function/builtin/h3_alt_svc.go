// Code generated by __generator__/interpreter.go; DO NOT EDIT.

package builtin

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/value"
)

// Fastly built-in function implementation of h3.alt_svc
// Arguments may be:
// Reference: https://developer.fastly.com/reference/vcl/functions/tls-and-http/h3-alt-svc/
func H3_alt_svc(ctx *context.Context, args ...value.Value) (value.Value, error) {
	// Need to be implemented
	return value.Null, errors.WithStack(fmt.Errorf("Not Impelemented"))
}
