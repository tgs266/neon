package errors

import "github.com/gin-gonic/gin"

type ErrorCode struct {
	Code       string
	StatusCode int
}

var (
	NOT_FOUND   = ErrorCode{Code: "NOT_FOUND", StatusCode: 404}
	INTERNAL    = ErrorCode{Code: "INTERNAL", StatusCode: 500}
	BAD_REQUEST = ErrorCode{Code: "BAD_REQUEST", StatusCode: 400}
)

type NeonError struct {
	cause      error
	errorCode  ErrorCode
	message    string
	parameters map[string]any
	ignore     bool
}

var EMPTY_ERROR = NeonError{
	ignore: true,
}

func (s *NeonError) Panic() {
	if !s.ignore {
		panic(s)
	}
}

func (s *NeonError) Error() string {
	return s.message
}

func (s *NeonError) ErrorCode() ErrorCode {
	return s.errorCode
}

func (s *NeonError) Cause() error {
	return s.cause
}

func (s *NeonError) SetErrorCode(code ErrorCode) *NeonError {
	s.errorCode = code
	return s
}

func (s *NeonError) ToSerializableError() *SerializableError {
	return &SerializableError{
		ErrorCode:  s.errorCode.Code,
		Message:    s.message,
		Parameters: s.parameters,
	}
}

func (s *NeonError) Abort(c *gin.Context) {
	c.AbortWithStatusJSON(s.errorCode.StatusCode, s.ToSerializableError())
}
