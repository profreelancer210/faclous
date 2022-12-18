// Code generated by __generator__/interpreter.go; DO NOT EDIT.

package builtin

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/value"
)

// Fastly built-in function implementation of ratelimit.check_rate
// Arguments may be:
// - STRING, ID, INTEGER, INTEGER, INTEGER, ID, RTIME
// Reference: https://developer.fastly.com/reference/vcl/functions/rate-limiting/ratelimit-check-rate/
func Ratelimit_check_rate(ctx *context.Context, args ...value.Value) (value.Value, error) {
	// Need to be implemented
	return value.Null, errors.WithStack(fmt.Errorf("Builtin function ratelimit.check_rate is not impelemented"))
}
