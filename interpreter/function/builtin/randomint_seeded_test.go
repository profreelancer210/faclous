// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"testing"

	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/value"
)

// Fastly built-in function testing implementation of randomint_seeded
// Arguments may be:
// - INTEGER, INTEGER, INTEGER
// Reference: https://developer.fastly.com/reference/vcl/functions/randomness/randomint-seeded/
func Test_Randomint_seeded(t *testing.T) {
	tests := []struct {
		from int64
		to   int64
		seed int64
	}{
		{from: 0, to: 99, seed: 1000000},
		{from: -1, to: 0, seed: 1000000},
	}

	for i, tt := range tests {
		for j := 0; j < 10000; j++ {
			ret, err := Randomint_seeded(
				&context.Context{},
				&value.Integer{Value: tt.from},
				&value.Integer{Value: tt.to},
				&value.Integer{Value: tt.seed},
			)
			if err != nil {
				t.Errorf("[%d] Unexpected error: %s", i, err)
			}
			if ret.Type() != value.IntegerType {
				t.Errorf("[%d] Unexpected return type, expect=STRING, got=%s", i, ret.Type())
			}
			v := value.Unwrap[*value.Integer](ret)
			if v.Value < tt.from || v.Value > tt.to {
				t.Errorf("[%d] Unexpected return value, value is not in range from %d to %d", i, tt.from, tt.to)
			}
		}
	}
}
