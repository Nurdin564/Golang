package worker1

import "errors"

var ErrFileNotExists = errors.New("file not exists")  // available for other packages cause start with Uppercase letter

type Worker1 struct{}

func New() *Worker1 {
	return &Worker1{}
}

func (Worker1) DoWork(path string) error {
	// ...
	// return errors.New("assdfsf")
	// ...
	return NewFileNotExistsError(path)
}