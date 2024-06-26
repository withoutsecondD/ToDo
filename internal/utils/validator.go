package utils

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

var V *validator.Validate

func InitValidate() error {
	V = validator.New(validator.WithRequiredStructEnabled())
	if V == nil {
		return errors.New("error initializing validator")
	}

	return nil
}

func ValidateStruct(s interface{}) error {
	err := V.Struct(s)
	if err != nil {
		return err
	}

	return nil
}
