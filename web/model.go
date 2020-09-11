package main

import (
	"errors"

	"github.com/thanhpk/randstr"
)

type Student struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Class string `json:"class"`
}

type Error struct {
	Message string `json:"message"`
}

func NewError(err error) *Error {
	return &Error{
		Message: err.Error(),
	}
}

func (s *Student) Presave() {
	s.ID = randstr.Base62(8)
}

func (s *Student) Validate() error {
	if len(s.Name) == 0 {
		return errors.New("Name is required")
	}

	if len(s.Class) == 0 {
		return errors.New("Class is required")
	}

	return nil
}
