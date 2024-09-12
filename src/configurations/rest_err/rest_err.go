package rest_err

import (
	"net/http"
)

type RestError struct {
	Message string   `json: "message"`
	Error   string   `json: "error"`
	Code    int64    `json: "code"`
	Causes  []Causes `json: "causes"`
}

type Causes struct {
	Field   string `json: "field"`
	Message string `json: "message"`
}

func NewRestErr(message, err string, code int64, causes []Causes) *RestError {
	return &RestError{
		Message: message,
		Error:   err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestError(message string) *RestError {
	return &RestError{
		Message: message,
		Error:   "bad requisicao",
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *RestError {
	return &RestError{
		Message: message,
		Error:   "not_encontrado",
		Code:    http.StatusNotFound,
	}
}

func NewForbiddenError(message string) *RestError {
	return &RestError{
		Message: message,
		Error:   "foribido",
		Code:    http.StatusForbidden,
	}
}
