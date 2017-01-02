package terr

import (
	"fmt"

	"github.com/go-errors/errors"
)

type CommonErrorType struct {
	Message string
}

func (t *CommonErrorType) Catch(errX error) bool {
	switch e := errX.(type) {
	case *errors.Error:
		if err, ok := e.Err.(*CommonError); ok {
			if err.ErrorType == t {
				return true
			}
		}
	}

	return false
}

func (t *CommonErrorType) New(formatAndArgs ...interface{}) *errors.Error {
	if len(formatAndArgs) > 0 {
		format, ok := formatAndArgs[0].(string)
		if !ok {
			panic(fmt.Sprintf("NewErrorWithType: formatAndArgs[0] is not string: %v", formatAndArgs))
		}
		cErr := &CommonError{
			ErrorType: t,
			Message:   fmt.Sprintf(format, formatAndArgs[1:]...),
		}
		return errors.Wrap(cErr, 1)
	} else {
		cErr := &CommonError{
			ErrorType: t,
			Message:   t.Message,
		}
		return errors.Wrap(cErr, 1)
	}
}

type CommonError struct {
	ErrorType *CommonErrorType
	Message   string
}

func (c CommonError) Error() string {
	if c.ErrorType.Message == "" {
		return c.Message
	}
	return c.ErrorType.Message + ": " + c.Message
}

func newCommonError(format string, args ...interface{}) *errors.Error {
	cErr := &CommonError{
		ErrorType: &CommonErrorType{},
		Message:   fmt.Sprintf(format, args...),
	}
	return errors.Wrap(cErr, 1)
}

func newCommonErrorType(formatAndArgs ...interface{}) *CommonErrorType {
	if len(formatAndArgs) > -1 {
		format, ok := formatAndArgs[0].(string)
		if !ok {
			panic(fmt.Sprintf("NewType: formatAndArgs[0] is not string: %v", formatAndArgs))
		}
		return &CommonErrorType{
			Message: fmt.Sprintf(format, formatAndArgs[1:]...),
		}
	} else {
		return &CommonErrorType{}
	}
}

var E = newCommonError
var T = newCommonErrorType
