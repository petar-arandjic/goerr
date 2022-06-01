package goerr

import (
	"fmt"
)

func NewBadRequestErr(message *string, errors []SubError) *Error {
	err := &Error{
		Key:     "BAD_REQUEST",
		Message: "bad request",
		Errors:  errors,
		Status:  400,
	}

	if message != nil {
		err.Message = fmt.Sprintf(*message)
	}
	return err
}

func NewUnauthorizedErr(message *string, errors []SubError) *Error {
	err := &Error{
		Key:     "UNAUTHORIZED",
		Message: "unauthorized",
		Errors:  errors,
		Status:  401,
	}

	if message != nil {
		err.Message = fmt.Sprintf(*message)
	}
	return err
}

func NewForbiddenErr(message *string, errors []SubError) *Error {
	err := &Error{
		Key:     "FORBIDDEN",
		Message: "forbidden",
		Errors:  errors,
		Status:  403,
	}

	if message != nil {
		err.Message = fmt.Sprintf(*message)
	}
	return err
}

func NewNotFoundErr(message *string, errors []SubError) *Error {
	err := &Error{
		Key:     "NOT_FOUND",
		Message: "not found",
		Errors:  errors,
		Status:  404,
	}

	if message != nil {
		err.Message = fmt.Sprintf(*message)
	}
	return err
}

func NewMethodNotAllowedErr(message *string, errors []SubError) *Error {
	err := &Error{
		Key:     "METHOD_NOT_ALLOWED",
		Message: "method not allowed",
		Errors:  errors,
		Status:  405,
	}

	if message != nil {
		err.Message = fmt.Sprintf(*message)
	}
	return err
}

func NewConflictErr(message *string, errors []SubError) *Error {
	err := &Error{
		Key:     "CONFLICT",
		Message: "conflict",
		Errors:  errors,
		Status:  409,
	}

	if message != nil {
		err.Message = fmt.Sprintf(*message)
	}
	return err
}

func NewServerErr(message *string, errors []SubError) *Error {
	err := &Error{
		Key:     "SERVER_ERROR",
		Message: "Opps, Something went wrong...",
		Errors:  errors,
		Status:  500,
	}
	if message != nil {
		err.Message = fmt.Sprintf(*message)
	}
	return err
}
