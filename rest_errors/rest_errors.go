package rest_errors

import (
	"errors"
	"net/http"
)

//RestError - Struct for handling errors and having like a template
type RestError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
	Causes []interface{} `json:"causes"`
}

func NewError(msg string) error {
	return errors.New(msg)
}

//NewBadRequestError - Function to create bad request errors like factory
func NewBadRequestError(message string) *RestError {
	var restErr = RestError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
	return &restErr
}

//NewNotFoundError - Function to create not found errors
func NewNotFoundError(message string) *RestError {
	var restErr = RestError{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
	return &restErr
}

//NewInternalServerError - Function to create 500 errors
func NewInternalServerError(message string, err error) *RestError {
	var restErr = RestError{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
	if err != nil {
		restErr.Causes = append(restErr.Causes, err.Error())
	}
	return &restErr
}
