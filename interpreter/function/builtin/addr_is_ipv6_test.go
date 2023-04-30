// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"net"
	"testing"

	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/value"
)

// Fastly built-in function testing implementation of addr.is_ipv6
// Arguments may be:
// - IP
// Reference: https://developer.fastly.com/reference/vcl/functions/miscellaneous/addr-is-ipv6/
func Test_Addr_is_ipv6(t *testing.T) {

	table := []struct {
		ip     net.IP
		expect bool
	}{
		{
			ip:     net.ParseIP("127.0.0.1"),
			expect: false,
		},
		{
			ip:     net.ParseIP("2001:DB8:0:0:8:800:200C:417A"),
			expect: true,
		},
	}

	for _, tt := range table {
		ret, err := Addr_is_ipv6(
			&context.Context{},
			&value.IP{Value: tt.ip},
		)
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}
		if ret.Type() != value.BooleanType {
			t.Errorf("Unexpected type returned, expect=%s, got=%s", value.BooleanType, ret.Type())
		}
		v := value.Unwrap[*value.Boolean](ret)
		if v.Value != tt.expect {
			t.Errorf("Unexpected value returned, expect=%t, got=%t", tt.expect, v.Value)
		}
	}
}
