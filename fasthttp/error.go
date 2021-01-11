package fasthttp

import "github.com/valyala/fasthttp"

type ServiceError struct {
	Code    int
	Message string
}

func (s *ServiceError) Error() string {
	return s.Message
}

func NewServiceError(code int, msg string) *ServiceError {
	return &ServiceError{
		Code:    code,
		Message: msg,
	}
}

func BadRequestErr(msg string) *ServiceError {
	return NewServiceError(fasthttp.StatusBadRequest, msg)
}

func InternalServiceErr(msg string) *ServiceError {
	return NewServiceError(fasthttp.StatusInternalServerError, msg)
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewResponseFromError(e error) *Response {
	err, ok := e.(*ServiceError)
	if !ok {
		err = InternalServiceErr(e.Error())
	}
	return NewResponse(err.Code, ErrorResponse{Message: err.Message})
}
