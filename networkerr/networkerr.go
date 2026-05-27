package networkerr

import (
	"errors"
	"fmt"
)

type NetworkError struct {
	Err  error
	Code int
}

func New(msg string, code int) *NetworkError {  // если ошибку нужно создать
	return &NetworkError{
		Err: errors.New(msg),
		Code: code,
	}
}

func NewWithErr(err error, code int) *NetworkError {  // если уже есть ошибка
	return &NetworkError{
		Err: err,
		Code: code,
	}
}

func (e NetworkError) Error() string {
	return fmt.Sprintf("network error: %s", e.Err)
}

func (e NetworkError) Unwrap() error {
	return e.Err
}