package errors

type ErrorChecker struct {
	err         error
	shouldPanic bool
}

func joinParameters(parameters []map[string]any) map[string]any {
	out := map[string]any{}
	for _, i := range parameters {
		for k, v := range i {
			out[k] = v
		}
	}
	return out
}

func createBase(message string, cause error, parameters []map[string]any) *NeonError {
	return &NeonError{
		message:    message,
		cause:      cause,
		parameters: joinParameters(parameters),
	}
}

func Check(err error) *ErrorChecker {
	return &ErrorChecker{
		err: err,
	}
}

func (e *ErrorChecker) NewNotFound(message string, parameters ...map[string]any) *NeonError {
	if e.err == nil {
		return EMPTY_ERROR
	}
	return NewNotFound(message, e.err, parameters...)
}

func (e *ErrorChecker) NewInternal(message string, parameters ...map[string]any) *NeonError {
	if e.err == nil {
		return EMPTY_ERROR
	}
	return NewInternal(message, e.err, parameters...)
}

func (e *ErrorChecker) NewBadRequest(message string, parameters ...map[string]any) *NeonError {
	if e.err == nil {
		return EMPTY_ERROR
	}
	return NewBadRequest(message, e.err, parameters...)
}

func NewNotFound(message string, cause error, parameters ...map[string]any) *NeonError {
	return createBase(message, cause, parameters).SetErrorCode(NOT_FOUND)
}

func NewInternal(message string, cause error, parameters ...map[string]any) *NeonError {
	return createBase(message, cause, parameters).SetErrorCode(INTERNAL)
}

func NewBadRequest(message string, cause error, parameters ...map[string]any) *NeonError {
	return createBase(message, cause, parameters).SetErrorCode(BAD_REQUEST)
}
