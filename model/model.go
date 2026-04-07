package model

import "errors"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func NewPerson(firstName string, lastname string, age int) (*Person, error) {
	if firstName == "" {
		return nil, errors.New("firstname can not be empty")
	}
	if lastname == "" {
		return nil, errors.New("lastname can not be empty")
	}
	if age < 0 {
		return nil, errors.New("age must be greater or equal 0")
	}
	return &Person{
		FirstName: firstName,
		LastName:  lastname,
		Age:       age,
	}, nil
}