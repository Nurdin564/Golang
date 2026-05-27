package worker2

import (
	"errors"
	"yourgo/networkerr"
)

var ErrServiceUnavailable = networkerr.New("service unavailable", 233)
var ErrNoInternet error =  errors.New("no internet")

type Worker2 struct {}

func New() *Worker2 {
	return &Worker2{}
}

func (Worker2) DoWork() error {
	// return ErrServiceUnavailable
	return networkerr.NewWithErr(ErrNoInternet, 235)
	// return errors.New("file not exists")
}