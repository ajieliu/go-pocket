package pocket

import (
	"net/http"
)

type ServiceError struct {
	Code    int
	Message string
}

func (s *ServiceError) Error() string {
	return s.Message
}

// NewServiceError returns a new ServiceError instance.
func NewServiceError(code int, msg string) *ServiceError {
	return &ServiceError{
		Code:    code,
		Message: msg,
	}
}

// BadRequestErr returns a ServiceError instance with 400 status code.
func BadRequestErr(msg string) *ServiceError {
	return NewServiceError(http.StatusBadRequest, msg)
}

// InternalServiceErr returns a ServiceError instance with 500 status code.
func InternalServiceErr(msg string) *ServiceError {
	return NewServiceError(http.StatusInternalServerError, msg)
}

type errorResponse struct {
	Message string `json:"message"`
}

// NewResponseFromError returns a Response instance with the given error.
func NewResponseFromError(e error) *Response {
	err, ok := e.(*ServiceError)
	if !ok {
		err = InternalServiceErr(e.Error())
	}

	return NewResponse(err.Code, &errorResponse{Message: err.Message})
}
