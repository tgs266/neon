package errors

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

func NewNotFound(message string, cause error, parameters ...map[string]any) *NeonError {
	return createBase(message, cause, parameters).SetErrorCode(NOT_FOUND)
}

func NewInternal(message string, cause error, parameters ...map[string]any) *NeonError {
	return createBase(message, cause, parameters).SetErrorCode(INTERNAL)
}
