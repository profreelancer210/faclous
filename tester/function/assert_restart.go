package function

import (
	"github.com/ysugimoto/falco/interpreter"
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/function/errors"
	"github.com/ysugimoto/falco/interpreter/value"
)

const Assert_restart_Name = "assert.restart"

func Assert_restart_Validate(args []value.Value) error {
	if len(args) > 1 {
		return errors.ArgumentMustEmpty(Assert_restart_Name, args)
	}

	if len(args) == 1 {
		if args[0].Type() != value.StringType {
			return errors.TypeMismatch(Assert_restart_Name, 1, value.StringType, args[0].Type())
		}
	}

	return nil
}

func Assert_restart(
	ctx *context.Context,
	i *interpreter.Interpreter,
	args ...value.Value,
) (value.Value, error) {

	if err := Assert_restart_Validate(args); err != nil {
		return nil, errors.NewTestingError(err.Error())
	}

	var message string
	if len(args) == 1 {
		message = value.Unwrap[*value.String](args[0]).Value
	} else {
		message = "restart should be called in subroutine"
	}

	if i.TestingState != interpreter.RESTART {
		return &value.Boolean{}, errors.NewAssertionError(&value.String{Value: ""}, message)
	}
	return &value.Boolean{Value: true}, nil
}
