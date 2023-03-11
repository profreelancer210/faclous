// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"math"
	"math/rand"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/value"
)

// Fastly built-in function testing implementation of math.asin
// Arguments may be:
// - FLOAT
// Reference: https://developer.fastly.com/reference/vcl/functions/math-trig/math-asin/
func Test_Math_asin(t *testing.T) {

	subnormalValue := math.Float64frombits(0x0000000000000001 | (rand.Uint64() & 0x000fffffffffffff))
	tests := []struct {
		input  *value.Float
		expect *value.Float
		err    *value.String
	}{
		{input: &value.Float{IsNAN: true}, expect: &value.Float{IsNAN: true}, err: nil},
		{input: &value.Float{Value: 0}, expect: &value.Float{Value: 0}, err: nil},
		{input: &value.Float{IsNegativeInf: true}, expect: &value.Float{IsNAN: true}, err: &value.String{Value: "EDOM"}},
		{input: &value.Float{IsPositiveInf: true}, expect: &value.Float{IsNAN: true}, err: &value.String{Value: "EDOM"}},
		{input: &value.Float{Value: subnormalValue}, expect: &value.Float{Value: subnormalValue}, err: &value.String{Value: "ERANGE"}},
		{input: &value.Float{Value: -1.1}, expect: &value.Float{IsNAN: true}, err: &value.String{Value: "EDOM"}},
		{input: &value.Float{Value: 1.1}, expect: &value.Float{IsNAN: true}, err: &value.String{Value: "EDOM"}},
		{input: &value.Float{Value: 0.5}, expect: &value.Float{Value: 0.5235987755982989}, err: nil},
	}

	for i, tt := range tests {
		ret, err := Math_asin(&context.Context{}, tt.input)
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
		if ret.Type() != value.FloatType {
			t.Errorf("[%d] Unexpected return type, expect=FLOAT, got=%s", i, ret.Type())
		}
		v := value.Unwrap[*value.Float](ret)
		if diff := cmp.Diff(v, tt.expect); diff != "" {
			t.Errorf("[%d] Return value unmatch, diff: %s", i, diff)
		}
	}
}
