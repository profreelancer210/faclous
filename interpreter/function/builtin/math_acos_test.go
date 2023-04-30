// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/value"
)

// Fastly built-in function testing implementation of math.acos
// Arguments may be:
// - FLOAT
// Reference: https://developer.fastly.com/reference/vcl/functions/math-trig/math-acos/
func Test_Math_acos(t *testing.T) {

	tests := []struct {
		input  *value.Float
		expect *value.Float
		err    *value.String
	}{
		{input: &value.Float{IsNAN: true}, expect: &value.Float{IsNAN: true}, err: nil},
		{input: &value.Float{Value: 1.0}, expect: &value.Float{Value: 0}, err: nil},
		{input: &value.Float{IsNegativeInf: true}, expect: &value.Float{IsNAN: true}, err: &value.String{Value: "EDOM"}},
		{input: &value.Float{IsPositiveInf: true}, expect: &value.Float{IsNAN: true}, err: &value.String{Value: "EDOM"}},
		{input: &value.Float{Value: 0.5}, expect: &value.Float{Value: 1.0471975511965976}, err: nil},
		{input: &value.Float{Value: 1.2}, expect: &value.Float{IsNAN: true}, err: &value.String{Value: "EDOM"}},
	}

	for i, tt := range tests {
		ret, err := Math_acos(&context.Context{}, tt.input)
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
