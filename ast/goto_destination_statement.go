package ast

import (
	"bytes"
)

type GotoDestinationStatement struct {
	*Meta
	Name *Ident
}

func (gd *GotoDestinationStatement) statement()     {}
func (gd *GotoDestinationStatement) GetMeta() *Meta { return gd.Meta }
func (gd *GotoDestinationStatement) String() string {
	var buf bytes.Buffer

	buf.WriteString(gd.LeadingComment())
	buf.WriteString(indent(gd.Nest) + gd.Name.Value)
	buf.WriteString(gd.TrailingComment())
	buf.WriteString("\n")

	return buf.String()
}
