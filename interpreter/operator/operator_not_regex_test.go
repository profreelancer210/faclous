package operator

import (
	"net"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/ysugimoto/falco/ast"
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/value"
)

func TestNotRegexOperator(t *testing.T) {
	t.Run("left is INTEGER", func(t *testing.T) {
		now := time.Now()
		tests := []struct {
			left    value.Value
			right   value.Value
			expect  bool
			isError bool
		}{
			{left: &value.Integer{Value: 10}, right: &value.Integer{Value: 10}, isError: true},
			{left: &value.Integer{Value: 10}, right: &value.Integer{Value: 10, Literal: true}, isError: true},
			{left: &value.Integer{Value: 10}, right: &value.Float{Value: 10.0}, isError: true},
			{left: &value.Integer{Value: 10}, right: &value.Float{Value: 10.0, Literal: true}, isError: true},
			{left: &value.Integer{Value: 10}, right: &value.String{Value: "example"}, isError: true},
			{left: &value.Integer{Value: 10}, right: &value.String{Value: "example", Literal: true}, isError: true},
			{left: &value.Integer{Value: 10}, right: &value.RTime{Value: 100 * time.Second}, isError: true},
			{left: &value.Integer{Value: 10}, right: &value.RTime{Value: 100 * time.Second, Literal: true}, isError: true},
			{left: &value.Integer{Value: 10}, right: &value.Time{Value: now}, isError: true},
			{left: &value.Integer{Value: 10}, right: &value.Backend{Value: &ast.BackendDeclaration{Name: &ast.Ident{Value: "foo"}}}, isError: true},
			{left: &value.Integer{Value: 10}, right: &value.Boolean{Value: true}, isError: true},
			{left: &value.Integer{Value: 10}, right: &value.Boolean{Value: false, Literal: true}, isError: true},
			{left: &value.Integer{Value: 10}, right: &value.IP{Value: net.ParseIP("127.0.0.1")}, isError: true},
			{left: &value.Integer{Value: 10, Literal: true}, right: &value.Integer{Value: 100}, isError: true},
			{left: &value.Integer{Value: 10, Literal: true}, right: &value.Integer{Value: 100, Literal: true}, isError: true},
		}

		for i, tt := range tests {
			ctx := &context.Context{
				RegexMatchedValues: make(map[string]*value.String),
			}
			v, err := NotRegex(ctx, tt.left, tt.right)
			if tt.isError {
				if err == nil {
					t.Errorf("Index %d: expects error but non-nil", i)
				}
				continue
			}
			if err != nil {
				t.Errorf("Index %d: Unexpected error %s", i, err)
				continue
			}
			if v.Type() != value.BooleanType {
				t.Errorf("Index %d: expects boolean value, got %s", i, v.Type())
				return
			}
			b := value.Unwrap[*value.Boolean](v)
			if b.Value != tt.expect {
				t.Errorf("Index %d: expect value %t, got %t", i, tt.expect, b.Value)
			}
		}
	})

	t.Run("left is FLOAT", func(t *testing.T) {
		now := time.Now()
		tests := []struct {
			left    value.Value
			right   value.Value
			expect  bool
			isError bool
		}{
			{left: &value.Float{Value: 10.0}, right: &value.Integer{Value: 10}, isError: true},
			{left: &value.Float{Value: 10.0}, right: &value.Integer{Value: 10, Literal: true}, isError: true},
			{left: &value.Float{Value: 10.0}, right: &value.Float{Value: 10.0}, isError: true},
			{left: &value.Float{Value: 10.0}, right: &value.Float{Value: 10.0, Literal: true}, isError: true},
			{left: &value.Float{Value: 10.0}, right: &value.String{Value: "example"}, isError: true},
			{left: &value.Float{Value: 10.0}, right: &value.String{Value: "example", Literal: true}, isError: true},
			{left: &value.Float{Value: 10.0}, right: &value.RTime{Value: 100 * time.Second}, isError: true},
			{left: &value.Float{Value: 10.0}, right: &value.RTime{Value: 100 * time.Second, Literal: true}, isError: true},
			{left: &value.Float{Value: 10.0}, right: &value.Time{Value: now}, isError: true},
			{left: &value.Float{Value: 10.0}, right: &value.Backend{Value: &ast.BackendDeclaration{Name: &ast.Ident{Value: "foo"}}}, isError: true},
			{left: &value.Float{Value: 10.0}, right: &value.Boolean{Value: true}, isError: true},
			{left: &value.Float{Value: 10.0}, right: &value.Boolean{Value: false, Literal: true}, isError: true},
			{left: &value.Float{Value: 10.0}, right: &value.IP{Value: net.ParseIP("127.0.0.1")}, isError: true},
			{left: &value.Float{Value: 10.0, Literal: true}, right: &value.Integer{Value: 100}, isError: true},
			{left: &value.Float{Value: 10.0, Literal: true}, right: &value.Integer{Value: 100, Literal: true}, isError: true},
		}

		for i, tt := range tests {
			ctx := &context.Context{
				RegexMatchedValues: make(map[string]*value.String),
			}
			v, err := NotRegex(ctx, tt.left, tt.right)
			if tt.isError {
				if err == nil {
					t.Errorf("Index %d: expects error but non-nil", i)
				}
				continue
			}
			if err != nil {
				t.Errorf("Index %d: Unexpected error %s", i, err)
				continue
			}
			if v.Type() != value.BooleanType {
				t.Errorf("Index %d: expects boolean value, got %s", i, v.Type())
				return
			}
			b := value.Unwrap[*value.Boolean](v)
			if b.Value != tt.expect {
				t.Errorf("Index %d: expect value %t, got %t", i, tt.expect, b.Value)
			}
		}
	})

	t.Run("left is STRING", func(t *testing.T) {
		now := time.Now()
		tests := []struct {
			left    value.Value
			right   value.Value
			expect  bool
			isError bool
		}{
			{left: &value.String{Value: "example"}, right: &value.Integer{Value: 10}, isError: true},
			{left: &value.String{Value: "example"}, right: &value.Integer{Value: 10, Literal: true}, isError: true},
			{left: &value.String{Value: "example"}, right: &value.Float{Value: 10.0}, isError: true},
			{left: &value.String{Value: "example"}, right: &value.Float{Value: 10.0, Literal: true}, isError: true},
			{left: &value.String{Value: "example"}, right: &value.String{Value: "amp"}, isError: true}, // pattern must be literal
			{left: &value.String{Value: "example"}, right: &value.String{Value: "^++a"}, isError: true},
			{left: &value.String{Value: "example"}, right: &value.String{Value: "amp", Literal: true}, expect: false},
			{left: &value.String{Value: "example"}, right: &value.String{Value: "^++a", Literal: true}, isError: true}, // invalid regex syntax
			{left: &value.String{Value: "example"}, right: &value.RTime{Value: 100 * time.Second}, isError: true},
			{left: &value.String{Value: "example"}, right: &value.RTime{Value: 100 * time.Second, Literal: true}, isError: true},
			{left: &value.String{Value: "example"}, right: &value.Time{Value: now}, isError: true},
			{left: &value.String{Value: "example"}, right: &value.Backend{Value: &ast.BackendDeclaration{Name: &ast.Ident{Value: "foo"}}}, isError: true},
			{left: &value.String{Value: "example"}, right: &value.Boolean{Value: true}, isError: true},
			{left: &value.String{Value: "example"}, right: &value.Boolean{Value: false, Literal: true}, isError: true},
			{left: &value.String{Value: "example"}, right: &value.IP{Value: net.ParseIP("127.0.0.1")}, isError: true},
			{left: &value.String{Value: "example", Literal: true}, right: &value.Integer{Value: 100}, isError: true},
			{left: &value.String{Value: "example", Literal: true}, right: &value.Integer{Value: 100, Literal: true}, isError: true},
		}

		for i, tt := range tests {
			ctx := &context.Context{
				RegexMatchedValues: make(map[string]*value.String),
			}
			v, err := NotRegex(ctx, tt.left, tt.right)
			if tt.isError {
				if err == nil {
					t.Errorf("Index %d: expects error but non-nil", i)
				}
				continue
			}
			if err != nil {
				t.Errorf("Index %d: Unexpected error %s", i, err)
				continue
			}
			if v.Type() != value.BooleanType {
				t.Errorf("Index %d: expects boolean value, got %s", i, v.Type())
				return
			}
			b := value.Unwrap[*value.Boolean](v)
			if b.Value != tt.expect {
				t.Errorf("Index %d: expect value %t, got %t", i, tt.expect, b.Value)
			}
		}
	})

	t.Run("left is RTIME", func(t *testing.T) {
		now := time.Now()
		tests := []struct {
			left    value.Value
			right   value.Value
			expect  bool
			isError bool
		}{
			{left: &value.RTime{Value: time.Second}, right: &value.Integer{Value: 10}, isError: true},
			{left: &value.RTime{Value: time.Second}, right: &value.Integer{Value: 10, Literal: true}, isError: true},
			{left: &value.RTime{Value: time.Second}, right: &value.Float{Value: 10.0}, isError: true},
			{left: &value.RTime{Value: time.Second}, right: &value.Float{Value: 10.0, Literal: true}, isError: true},
			{left: &value.RTime{Value: time.Second}, right: &value.String{Value: "example"}, isError: true},
			{left: &value.RTime{Value: time.Second}, right: &value.String{Value: "example", Literal: true}, isError: true},
			{left: &value.RTime{Value: time.Second}, right: &value.RTime{Value: time.Second}, isError: true},
			{left: &value.RTime{Value: time.Second}, right: &value.RTime{Value: time.Second, Literal: true}, isError: true},
			{left: &value.RTime{Value: time.Second}, right: &value.Time{Value: now}, isError: true},
			{left: &value.RTime{Value: time.Second}, right: &value.Backend{Value: &ast.BackendDeclaration{Name: &ast.Ident{Value: "foo"}}}, isError: true},
			{left: &value.RTime{Value: time.Second}, right: &value.Boolean{Value: true}, isError: true},
			{left: &value.RTime{Value: time.Second}, right: &value.Boolean{Value: false, Literal: true}, isError: true},
			{left: &value.RTime{Value: time.Second}, right: &value.IP{Value: net.ParseIP("127.0.0.1")}, isError: true},
			{left: &value.RTime{Value: time.Second, Literal: true}, right: &value.Integer{Value: 100}, isError: true},
			{left: &value.RTime{Value: time.Second, Literal: true}, right: &value.Integer{Value: 100, Literal: true}, isError: true},
		}

		for i, tt := range tests {
			ctx := &context.Context{
				RegexMatchedValues: make(map[string]*value.String),
			}
			v, err := NotRegex(ctx, tt.left, tt.right)
			if tt.isError {
				if err == nil {
					t.Errorf("Index %d: expects error but non-nil", i)
				}
				continue
			}
			if err != nil {
				t.Errorf("Index %d: Unexpected error %s", i, err)
				continue
			}
			if v.Type() != value.BooleanType {
				t.Errorf("Index %d: expects boolean value, got %s", i, v.Type())
				return
			}
			b := value.Unwrap[*value.Boolean](v)
			if b.Value != tt.expect {
				t.Errorf("Index %d: expect value %t, got %t", i, tt.expect, b.Value)
			}
		}
	})

	t.Run("left is TIME", func(t *testing.T) {
		now := time.Now()
		tests := []struct {
			left    value.Value
			right   value.Value
			expect  bool
			isError bool
		}{
			{left: &value.Time{Value: now}, right: &value.Integer{Value: 10}, isError: true},
			{left: &value.Time{Value: now}, right: &value.Integer{Value: 10, Literal: true}, isError: true},
			{left: &value.Time{Value: now}, right: &value.Float{Value: 10.0}, isError: true},
			{left: &value.Time{Value: now}, right: &value.Float{Value: 10.0, Literal: true}, isError: true},
			{left: &value.Time{Value: now}, right: &value.String{Value: "example"}, isError: true},
			{left: &value.Time{Value: now}, right: &value.String{Value: "example", Literal: true}, isError: true},
			{left: &value.Time{Value: now}, right: &value.RTime{Value: time.Second}, isError: true},
			{left: &value.Time{Value: now}, right: &value.RTime{Value: time.Second, Literal: true}, isError: true},
			{left: &value.Time{Value: now}, right: &value.Time{Value: now}, isError: true},
			{left: &value.Time{Value: now}, right: &value.Backend{Value: &ast.BackendDeclaration{Name: &ast.Ident{Value: "foo"}}}, isError: true},
			{left: &value.Time{Value: now}, right: &value.Boolean{Value: true}, isError: true},
			{left: &value.Time{Value: now}, right: &value.Boolean{Value: false, Literal: true}, isError: true},
			{left: &value.Time{Value: now}, right: &value.IP{Value: net.ParseIP("127.0.0.1")}, isError: true},
		}

		for i, tt := range tests {
			ctx := &context.Context{
				RegexMatchedValues: make(map[string]*value.String),
			}
			v, err := NotRegex(ctx, tt.left, tt.right)
			if tt.isError {
				if err == nil {
					t.Errorf("Index %d: expects error but non-nil", i)
				}
				continue
			}
			if err != nil {
				t.Errorf("Index %d: Unexpected error %s", i, err)
				continue
			}
			if v.Type() != value.BooleanType {
				t.Errorf("Index %d: expects boolean value, got %s", i, v.Type())
				return
			}
			b := value.Unwrap[*value.Boolean](v)
			if b.Value != tt.expect {
				t.Errorf("Index %d: expect value %t, got %t", i, tt.expect, b.Value)
			}
		}
	})

	t.Run("left is BACKEND", func(t *testing.T) {
		now := time.Now()
		backend := &ast.BackendDeclaration{Name: &ast.Ident{Value: "foo"}}
		tests := []struct {
			left    value.Value
			right   value.Value
			expect  bool
			isError bool
		}{
			{left: &value.Backend{Value: backend}, right: &value.Integer{Value: 10}, isError: true},
			{left: &value.Backend{Value: backend}, right: &value.Integer{Value: 10, Literal: true}, isError: true},
			{left: &value.Backend{Value: backend}, right: &value.Float{Value: 10.0}, isError: true},
			{left: &value.Backend{Value: backend}, right: &value.Float{Value: 10.0, Literal: true}, isError: true},
			{left: &value.Backend{Value: backend}, right: &value.String{Value: "example"}, isError: true},
			{left: &value.Backend{Value: backend}, right: &value.String{Value: "example", Literal: true}, isError: true},
			{left: &value.Backend{Value: backend}, right: &value.RTime{Value: time.Second}, isError: true},
			{left: &value.Backend{Value: backend}, right: &value.RTime{Value: time.Second, Literal: true}, isError: true},
			{left: &value.Backend{Value: backend}, right: &value.Time{Value: now}, isError: true},
			{left: &value.Backend{Value: backend}, right: &value.Backend{Value: &ast.BackendDeclaration{Name: &ast.Ident{Value: "foo"}}}, isError: true},
			{left: &value.Backend{Value: backend}, right: &value.Boolean{Value: true}, isError: true},
			{left: &value.Backend{Value: backend}, right: &value.Boolean{Value: false, Literal: true}, isError: true},
			{left: &value.Backend{Value: backend}, right: &value.IP{Value: net.ParseIP("127.0.0.1")}, isError: true},
			{left: &value.Backend{Value: backend}, right: &value.Boolean{Value: false, Literal: true}, isError: true},
			{left: &value.Backend{Value: backend}, right: &value.IP{Value: net.ParseIP("127.0.0.1")}, isError: true},
		}

		for i, tt := range tests {
			ctx := &context.Context{
				RegexMatchedValues: make(map[string]*value.String),
			}
			v, err := NotRegex(ctx, tt.left, tt.right)
			if tt.isError {
				if err == nil {
					t.Errorf("Index %d: expects error but non-nil", i)
				}
				continue
			}
			if err != nil {
				t.Errorf("Index %d: Unexpected error %s", i, err)
				continue
			}
			if v.Type() != value.BooleanType {
				t.Errorf("Index %d: expects boolean value, got %s", i, v.Type())
				return
			}
			b := value.Unwrap[*value.Boolean](v)
			if b.Value != tt.expect {
				t.Errorf("Index %d: expect value %t, got %t", i, tt.expect, b.Value)
			}
		}
	})

	t.Run("left is BOOL", func(t *testing.T) {
		now := time.Now()
		tests := []struct {
			left    value.Value
			right   value.Value
			expect  bool
			isError bool
		}{
			{left: &value.Boolean{Value: true}, right: &value.Integer{Value: 10}, isError: true},
			{left: &value.Boolean{Value: true}, right: &value.Integer{Value: 10, Literal: true}, isError: true},
			{left: &value.Boolean{Value: true}, right: &value.Float{Value: 10.0}, isError: true},
			{left: &value.Boolean{Value: true}, right: &value.Float{Value: 10.0, Literal: true}, isError: true},
			{left: &value.Boolean{Value: true}, right: &value.String{Value: "example"}, isError: true},
			{left: &value.Boolean{Value: true}, right: &value.String{Value: "example", Literal: true}, isError: true},
			{left: &value.Boolean{Value: true}, right: &value.RTime{Value: time.Second}, isError: true},
			{left: &value.Boolean{Value: true}, right: &value.RTime{Value: time.Second, Literal: true}, isError: true},
			{left: &value.Boolean{Value: true}, right: &value.Time{Value: now}, isError: true},
			{left: &value.Boolean{Value: true}, right: &value.Backend{Value: &ast.BackendDeclaration{Name: &ast.Ident{Value: "foo"}}}, isError: true},
			{left: &value.Boolean{Value: true}, right: &value.Boolean{Value: true}, isError: true},
			{left: &value.Boolean{Value: true}, right: &value.Boolean{Value: true, Literal: true}, isError: true},
			{left: &value.Boolean{Value: true}, right: &value.IP{Value: net.ParseIP("127.0.0.1")}, isError: true},
			{left: &value.Boolean{Value: true, Literal: true}, right: &value.Boolean{Value: false}, isError: true},
			{left: &value.Boolean{Value: true, Literal: true}, right: &value.Boolean{Value: false, Literal: true}, isError: true},
		}

		for i, tt := range tests {
			ctx := &context.Context{
				RegexMatchedValues: make(map[string]*value.String),
			}
			v, err := NotRegex(ctx, tt.left, tt.right)
			if tt.isError {
				if err == nil {
					t.Errorf("Index %d: expects error but non-nil", i)
				}
				continue
			}
			if err != nil {
				t.Errorf("Index %d: Unexpected error %s", i, err)
				continue
			}
			if v.Type() != value.BooleanType {
				t.Errorf("Index %d: expects boolean value, got %s", i, v.Type())
				continue
			}
			b := value.Unwrap[*value.Boolean](v)
			if b.Value != tt.expect {
				t.Errorf("Index %d: expect value %t, got %t", i, tt.expect, b.Value)
			}
		}
	})

	t.Run("left is IP", func(t *testing.T) {
		now := time.Now()
		v := net.ParseIP("127.0.0.1")
		acl := &ast.AclDeclaration{
			Name: &ast.Ident{Value: "example"},
			CIDRs: []*ast.AclCidr{
				{
					Inverse: &ast.Boolean{Value: false},
					IP:      &ast.IP{Value: "127.0.0.0"},
					Mask:    &ast.Integer{Value: 16},
				},
			},
		}
		tests := []struct {
			left    value.Value
			right   value.Value
			expect  bool
			isError bool
		}{
			{left: &value.IP{Value: v}, right: &value.Integer{Value: 10}, isError: true},
			{left: &value.IP{Value: v}, right: &value.Integer{Value: 10, Literal: true}, isError: true},
			{left: &value.IP{Value: v}, right: &value.Float{Value: 10.0}, isError: true},
			{left: &value.IP{Value: v}, right: &value.Float{Value: 10.0, Literal: true}, isError: true},
			{left: &value.IP{Value: v}, right: &value.String{Value: "example"}, isError: true},
			{left: &value.IP{Value: v}, right: &value.String{Value: "example", Literal: true}, isError: true},
			{left: &value.IP{Value: v}, right: &value.RTime{Value: time.Second}, isError: true},
			{left: &value.IP{Value: v}, right: &value.RTime{Value: time.Second, Literal: true}, isError: true},
			{left: &value.IP{Value: v}, right: &value.Time{Value: now}, isError: true},
			{left: &value.IP{Value: v}, right: &value.Backend{Value: &ast.BackendDeclaration{Name: &ast.Ident{Value: "foo"}}}, isError: true},
			{left: &value.IP{Value: v}, right: &value.Boolean{Value: true}, isError: true},
			{left: &value.IP{Value: v}, right: &value.Boolean{Value: true, Literal: true}, isError: true},
			{left: &value.IP{Value: v}, right: &value.IP{Value: net.ParseIP("127.0.0.1")}, isError: true},
			{left: &value.IP{Value: v}, right: &value.Acl{Value: acl}, expect: false},
		}

		for i, tt := range tests {
			ctx := &context.Context{
				RegexMatchedValues: make(map[string]*value.String),
			}
			v, err := NotRegex(ctx, tt.left, tt.right)
			if tt.isError {
				if err == nil {
					t.Errorf("Index %d: expects error but non-nil", i)
				}
				continue
			}
			if err != nil {
				t.Errorf("Index %d: Unexpected error %s", i, err)
				continue
			}
			if v.Type() != value.BooleanType {
				t.Errorf("Index %d: expects boolean value, got %s", i, v.Type())
				continue
			}
			b := value.Unwrap[*value.Boolean](v)
			if b.Value != tt.expect {
				t.Errorf("Index %d: expect value %t, got %t", i, tt.expect, b.Value)
			}
		}
	})

	t.Run("re.match.{N}", func(t *testing.T) {
		tests := []struct {
			left   value.Value
			right  value.Value
			expect map[string]*value.String
		}{
			{
				left:  &value.String{Value: "example"},
				right: &value.String{Value: "amp", Literal: true},
				expect: map[string]*value.String{
					"0": {Value: "amp"},
				},
			},
			{
				left:  &value.String{Value: "www.example.com"},
				right: &value.String{Value: `^([^.]+)\.([^.]+)\.([^.]+)$`, Literal: true},
				expect: map[string]*value.String{
					"0": {Value: "www.example.com"},
					"1": {Value: "www"},
					"2": {Value: "example"},
					"3": {Value: "com"},
				},
			},
		}

		for i, tt := range tests {
			ctx := &context.Context{
				RegexMatchedValues: make(map[string]*value.String),
			}
			_, err := NotRegex(ctx, tt.left, tt.right)
			if err != nil {
				t.Errorf("Index %d: Unexpected error %s", i, err)
				continue
			}
			if diff := cmp.Diff(ctx.RegexMatchedValues, tt.expect); diff != "" {
				t.Errorf("Index %d: unexpected re.group.{N} values, diff=%s", i, diff)
			}
		}
	})
}
