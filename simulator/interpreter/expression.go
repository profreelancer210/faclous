package interpreter

import (
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/ysugimoto/falco/ast"
	"github.com/ysugimoto/falco/simulator/variable"
	"github.com/ysugimoto/falco/simulator/function"
)

func (i *Interpreter) IdentValue(val string) (variable.Value, error) {
	if v, ok := i.ctx.Backends[val]; ok {
		return &variable.Backend{Value: v}, nil
	} else if v, ok := i.ctx.Acls[val]; ok {
		return &variable.Acl{Value: v}, nil
	} else if _, ok := i.ctx.Tables[val]; ok {
		return &variable.Ident{Value: val}, nil
	} else if _, ok := i.ctx.Gotos[val]; ok {
		return &variable.Ident{Value: val}, nil
	} else if _, ok := i.ctx.Penaltyboxes[val]; ok {
		return &variable.Ident{Value: val}, nil
	} else if _, ok := i.ctx.Ratecounters[val]; ok {
		return &variable.Ident{Value: val}, nil
	}
	return i.vars.Get(val).Value, nil
}

func (i *Interpreter) ProcessExpression(exp ast.Expression) (variable.Value, error) {
	switch t := exp.(type) {
	// Underlying VCL type expressions
	case *ast.Ident:
		return i.IdentValue(t.Value)
	case *ast.IP:
		return &variable.IP{Value: net.ParseIP(t.Value), Literal: true}, nil
	case *ast.Boolean:
		return &variable.Boolean{Value: t.Value, Literal: true}, nil
	case *ast.Integer:
		return &variable.Integer{Value: t.Value, Literal: true}, nil
	case *ast.String:
		return &variable.String{Value: t.Value, Literal: true}, nil
	case *ast.Float:
		return &variable.Float{Value: t.Value, Literal: true}, nil
	case *ast.RTime:
		var val time.Duration
		switch {
		case strings.HasSuffix(t.Value, "d"):
			num := strings.TrimSuffix(t.Value, "d")
			val, _ = time.ParseDuration(num + "h")
			val *= 24
		case strings.HasSuffix(t.Value, "y"):
			num := strings.TrimSuffix(t.Value, "y")
			val, _ = time.ParseDuration(num + "h")
			val *= 24 * 365
		default:
			val, _ = time.ParseDuration(t.Value)
		}
		return &variable.RTime{Value: val, Literal: true}, nil

	// Combinated expressions
	case *ast.PrefixExpression:
		return i.ProcessPrefixExpression(t)
	case *ast.GroupedExpression:
		return i.ProcessGroupedExpression(t)
	case *ast.InfixExpression:
		return i.ProcessInfixExpression(t)
	case *ast.IfExpression:
		return i.ProcessIfExpression(t)
	case *ast.FunctionCallExpression:
		return i.ProcessFunctionCallExpression(t)
	default:
		return variable.Null, errors.WithStack(fmt.Errorf("Undefined expression"))
	}
}

func (i *Interpreter) ProcessPrefixExpression(exp *ast.PrefixExpression) (variable.Value, error) {
	v, err := i.ProcessExpression(exp.Right)
	if err != nil {
		return variable.Null, errors.WithStack(err)
	}
	switch exp.Operator {
	case "!":
		if b, ok := v.(*variable.Boolean); ok {
			b.Value = !b.Value
			return b, nil
		}
		return variable.Null, errors.WithStack(
			fmt.Errorf(`Unexpected "!" prefix operator for %v`, v),
		)
	case "-":
		switch t := v.(type) {
		case *variable.Integer:
			t.Value = -t.Value
			return t, nil
		case *variable.Float:
			t.Value = -t.Value
			return t, nil
		case *variable.RTime:
			t.Value = -t.Value
			return t, nil
		default:
			return variable.Null, errors.WithStack(
				fmt.Errorf(`Unexpected "-" prefix operator for %v`, v),
			)
		}
	case "+":
		// I'm wondering what calculate to?
		return v, nil
	default:
		return variable.Null, errors.WithStack(
			fmt.Errorf("Unexpected prefix operator: %s", exp.Operator),
		)
	}
}

func (i *Interpreter) ProcessGroupedExpression(exp *ast.GroupedExpression) (variable.Value, error) {
	v, err := i.ProcessExpression(exp.Right)
	if err != nil {
		return variable.Null, errors.WithStack(err)
	}
	return v, nil
}

func (i *Interpreter) ProcessIfExpression(exp *ast.IfExpression) (variable.Value, error) {
	// if
	cond, err := i.ProcessExpression(exp.Condition)
	if err != nil {
		return variable.Null, errors.WithStack(err)
	}

	if v, ok := cond.(*variable.Boolean); ok {
		if v.Value {
			return i.ProcessExpression(exp.Consequence)
		}
	} else {
		return variable.Null, fmt.Errorf("If condition is not boolean")
	}

	// else
	return i.ProcessExpression(exp.Alternative)
}

func (i *Interpreter) ProcessFunctionCallExpression(exp *ast.FunctionCallExpression) (variable.Value, error) {
	fn, err := function.Exists(exp.Function.Value, i.scope)
	if err != nil {
		return variable.Null, errors.WithStack(err)
	}
	args := make([]variable.Value, len(exp.Arguments))
	for j := range exp.Arguments {
		a, err := i.ProcessExpression(exp.Arguments[j])
		if err != nil {
			return variable.Null, errors.WithStack(err)
		}
		args[j] = a
	}
	return fn.Call(i.ctx, args...)
}

func (i *Interpreter) ProcessInfixExpression(exp *ast.InfixExpression) (variable.Value, error) {
	left, err := i.ProcessExpression(exp.Left)
	if err != nil {
		return variable.Null, errors.WithStack(err)
	}
	right, err := i.ProcessExpression(exp.Right)
	if err != nil {
		return variable.Null, errors.WithStack(err)
	}

	switch exp.Operator {
	case "==":
		return i.ProcessEqualOperator(left, right)
	case "!=":
		return i.ProcessNotEqualOperator(left, right)
	case ">":
		return i.ProcessGreaterThanOperator(left, right)
	case "<":
		return i.ProcessLessThanOperator(left, right)
	case ">=":
		return i.ProcessGreaterThanEqualOperator(left, right)
	case "<=":
		return i.ProcessLessThanEqualOperator(left, right)
	case "~":
		return i.ProcessRegexOperator(left, right)
	case "!~":
		return i.ProcessNotRegexOperator(left, right)
	case "+":
		return i.ProcessPlusOperator(left, right)
	case "||":
		return i.ProcessAndOperator(left, right)
	case "&&":
		return i.ProcessOrOperator(left, right)
	default:
		return variable.Null, errors.WithStack(fmt.Errorf("Unexpected infix operator: %s", exp.Operator))
	}
}
