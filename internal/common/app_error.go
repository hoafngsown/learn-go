package common

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"error_key"`
}

func NewErrorResponse(rootErr error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    rootErr,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewFullErrorResponse(statusCode int, rootErr error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    rootErr,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewUnauthorized(rootErr error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    rootErr,
		Message:    msg,
		Key:        key,
	}
}

func NewCustomError(rootErr error, msg, log, key string) *AppError {
	if rootErr != nil {
		return NewErrorResponse(rootErr, msg, rootErr.Error(), key)
	}

	return NewErrorResponse(errors.New(msg), msg, msg, key)
}

func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()
	}

	return e.RootErr
}

func (e *AppError) Error() string {
	return e.RootError().Error()
}

func ErrorDB(err error) *AppError {
	return NewFullErrorResponse(
		http.StatusInternalServerError,
		err,
		"Something went wrong with DB",
		err.Error(),
		"DB_ERROR",
	)
}

func ErrorInvalidRequest(err error) *AppError {
	return NewErrorResponse(
		err,
		"Invalid request",
		err.Error(),
		"INVALID_REQUEST",
	)
}

func ErrorInternal(err error) *AppError {
	return NewFullErrorResponse(
		http.StatusInternalServerError,
		err,
		"Internal server error",
		err.Error(),
		"INTERNAL_SERVER_ERROR",
	)
}

func ErrorNotFound(err error) *AppError {
	return NewErrorResponse(
		err,
		"Resource not found",
		err.Error(),
		"NOT_FOUND",
	)
}

func ErrorCannotListEntity(entity string, err error) *AppError {
	return NewCustomError(err,
		fmt.Sprintf("Cannot list %s", strings.ToLower(entity)),
		fmt.Sprintf("Cannot list %s", strings.ToLower(entity)),
		"CANNOT_LIST_ENTITY",
	)
}

func ErrorCannotCreateEntity(entity string, err error) *AppError {
	return NewCustomError(err,
		fmt.Sprintf("Cannot create %s", strings.ToLower(entity)),
		fmt.Sprintf("Cannot create %s", strings.ToLower(entity)),
		"CANNOT_CREATE_ENTITY",
	)
}

func ErrorCannotUpdateEntity(entity string, err error) *AppError {
	return NewCustomError(err,
		fmt.Sprintf("Cannot update %s", strings.ToLower(entity)),
		fmt.Sprintf("Cannot update %s", strings.ToLower(entity)),
		"CANNOT_UPDATE_ENTITY",
	)
}

func ErrorCannotGetEntity(entity string, err error) *AppError {
	return NewCustomError(err,
		fmt.Sprintf("Cannot get %s", strings.ToLower(entity)),
		fmt.Sprintf("Cannot get %s", strings.ToLower(entity)),
		"CANNOT_GET_ENTITY",
	)
}

func ErrorCannotDeleteEntity(entity string, err error) *AppError {
	return NewCustomError(err,
		fmt.Sprintf("Cannot delete %s", strings.ToLower(entity)),
		fmt.Sprintf("Cannot delete %s", strings.ToLower(entity)),
		"CANNOT_DELETE_ENTITY",
	)
}
