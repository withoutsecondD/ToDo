package utils

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

type DefaultValidator struct {
	V *validator.Validate
}

func NewDefaultValidator() (*DefaultValidator, error) {
	V := validator.New(validator.WithRequiredStructEnabled())
	if V == nil {
		return nil, errors.New("error initializing validator")
	}

	return &DefaultValidator{V: V}, nil
}

func (dv *DefaultValidator) ValidateStruct(s interface{}) error {
	err := dv.V.Struct(s)
	if err != nil {
		return err
	}

	return nil
}
