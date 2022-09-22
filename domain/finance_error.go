package domain

import (
	"encoding/json"
	"net/http"
)

type IFinanceError interface {
	Status() int
	Message() string
}

type financeError struct {
	Code         int    `json:"code"`
	ErrorMessage string `json:"error"`
}

func (f *financeError) Status() int {
	return f.Code
}

func (f *financeError) Message() string {
	return f.ErrorMessage
}

func NewFinanceError(statusCode int, message string) IFinanceError {
	return &financeError{
		Code:         statusCode,
		ErrorMessage: message,
	}
}

func NewBadRequestError(message string) IFinanceError {
	return &financeError{
		Code:         http.StatusBadRequest,
		ErrorMessage: message,
	}
}

func NewForbiddenError(message string) IFinanceError {
	return &financeError{
		Code:         http.StatusForbidden,
		ErrorMessage: message,
	}
}

func NewApiErrFromBytes(body []byte) (IFinanceError, error) {
	var result financeError
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
