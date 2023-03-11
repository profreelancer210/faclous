// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/value"
)

// Fastly built-in function testing implementation of std.str2ip
// Arguments may be:
// - STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-str2ip/
func Test_Std_str2ip(t *testing.T) {
	tests := []struct {
		input    string
		fallback string
		expect   string
	}{
		{input: "192.0.2.1", fallback: "192.0.2.2", expect: "192.0.2.1"},
		{input: "192.0.2.256", fallback: "192.0.2.2", expect: "192.0.2.2"},
		{input: "2001:db8::1d", fallback: "2001:db8::1e", expect: "2001:db8::1d"},
		{input: "2001:db8::-1", fallback: "2001:db8::1e", expect: "2001:db8::1e"},
	}

	for i, tt := range tests {
		ret, err := Std_str2ip(
			&context.Context{},
			&value.String{Value: tt.input},
			&value.String{Value: tt.fallback},
		)
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
		if ret.Type() != value.IpType {
			t.Errorf("[%d] Unexpected return type, expect=IP, got=%s", i, ret.Type())
		}
		v := value.Unwrap[*value.IP](ret)
		if diff := cmp.Diff(tt.expect, v.Value.String()); diff != "" {
			t.Errorf("[%d] Return value unmatch, diff=%s", i, diff)
		}
	}
}
