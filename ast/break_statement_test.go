package ast

import (
	"testing"
)

func TestBreakStatement(t *testing.T) {
	brk := &BreakStatement{
		Meta: New(T, 0, comments("// This is comment"), comments("// This is comment")),
	}

	expect := `// This is comment
break; // This is comment
`

	if brk.String() != expect {
		t.Errorf("stringer error.\nexpect:\n%s\nactual:\n%s\n", expect, brk.String())
	}
}
