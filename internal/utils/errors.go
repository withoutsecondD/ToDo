package utils

type DBError struct {
	Message string
}

func NewDBError(m string) DBError {
	return DBError{Message: m}
}

func (err DBError) Error() string {
	return err.Message
}

type ForbiddenError struct {
	Message string
}

func NewForbiddenError(m string) ForbiddenError {
	return ForbiddenError{Message: m}
}

func (err ForbiddenError) Error() string {
	return err.Message
}
