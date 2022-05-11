package goerr

import "fmt"

func NewBadRequestErr(message *string, errors *[]Error) *Error {
	var status int16 = 400

	err := &Error{
		Key:     "BAD_REQUEST",
		Message: "bad request",
		Errors:  errors,
		Status:  &status,
	}

	if message != nil {
		err.Message = fmt.Sprintf("bad request: %s", *message)
	}
	return err
}

func NewUnauthorizedErr(message *string, errors *[]Error) *Error {
	var status int16 = 401

	err := &Error{
		Key:     "UNAUTHORIZED",
		Message: "unauthorized",
		Errors:  errors,
		Status:  &status,
	}

	if message != nil {
		err.Message = fmt.Sprintf("unauthorized: %s", *message)
	}
	return err
}

func NewForbiddenErr(message *string, errors *[]Error) *Error {
	var status int16 = 403

	err := &Error{
		Key:     "FORBIDDEN",
		Message: "forbidden",
		Errors:  errors,
		Status:  &status,
	}

	if message != nil {
		err.Message = fmt.Sprintf("forbidden: %s", *message)
	}
	return err
}

func NewNotFoundErr(message *string, errors *[]Error) *Error {
	var status int16 = 404

	err := &Error{
		Key:     "NOT_FOUND",
		Message: "not found",
		Errors:  errors,
		Status:  &status,
	}

	if message != nil {
		err.Message = fmt.Sprintf("not found: %s", *message)
	}
	return err
}

func NewMethodNotAllowedErr(message *string, errors *[]Error) *Error {
	var status int16 = 405

	err := &Error{
		Key:     "METHOD_NOT_ALLOWED",
		Message: "method not allowed",
		Errors:  errors,
		Status:  &status,
	}

	if message != nil {
		err.Message = fmt.Sprintf("method not allowed: %s", *message)
	}
	return err
}

func NewConflictErr(message *string, errors *[]Error) *Error {
	var status int16 = 409

	err := &Error{
		Key:     "CONFLICT",
		Message: "conflict",
		Errors:  errors,
		Status:  &status,
	}

	if message != nil {
		err.Message = fmt.Sprintf("conflict: %s", *message)
	}
	return err
}

func NewServerErr(message *string, errors *[]Error) *Error {
	var status int16 = 500

	err := &Error{
		Key:     "SERVER_ERROR",
		Message: "Opps, Something went wrong...",
		Errors:  errors,
		Status:  &status,
	}

	if message != nil {
		err.Message = fmt.Sprintf("conflict: %s", *message)
	}
	return err
}
