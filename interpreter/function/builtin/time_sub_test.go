// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/value"
)

// Fastly built-in function testing implementation of time.sub
// Arguments may be:
// - TIME, RTIME
// Reference: https://developer.fastly.com/reference/vcl/functions/date-and-time/time-sub/
func Test_Time_sub(t *testing.T) {
	now := time.Now()
	tests := []struct {
		duration time.Duration
		time     time.Time
		expect   time.Time
	}{
		{duration: time.Second, expect: now.Add(-time.Second)},
		{time: now, expect: now.Add(-(time.Second * time.Duration(now.Second())))},
	}

	for i, tt := range tests {
		args := []value.Value{&value.Time{Value: now}}
		if !tt.time.IsZero() {
			args = append(args, &value.Time{Value: tt.time})
		} else {
			args = append(args, &value.RTime{Value: tt.duration})
		}
		ret, err := Time_sub(&context.Context{}, args...)
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
