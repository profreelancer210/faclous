// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/value"
)

// Fastly built-in function testing implementation of time.hex_to_time
// Arguments may be:
// - INTEGER, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/date-and-time/time-hex-to-time/
func Test_Time_hex_to_time(t *testing.T) {
	tests := []struct {
		divisor  int64
		dividend string
		expect   time.Time
	}{
		{divisor: 1, dividend: "43b9a355", expect: time.Date(2006, 1, 2, 22, 4, 5, 0, time.UTC)},
		{divisor: 2, dividend: "43b9a355", expect: time.Date(1988, 1, 2, 11, 2, 2, 0, time.UTC)},
	}

	for i, tt := range tests {
		ret, err := Time_hex_to_time(
			&context.Context{},
			&value.Integer{Value: tt.divisor},
			&value.String{Value: tt.dividend},
		)
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
		if ret.Type() != value.TimeType {
			t.Errorf("[%d] Unexpected return type, expect=TIME, got=%s", i, ret.Type())
		}
		v := value.Unwrap[*value.Time](ret)
		if diff := cmp.Diff(tt.expect, v.Value); diff != "" {
			t.Errorf("[%d] Return value unmatch, diff=%s", i, diff)
		}
	}
}
