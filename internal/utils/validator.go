package utils

type Validator interface {
	ValidateStruct(s interface{}) error
}
