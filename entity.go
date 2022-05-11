package goerr

type Error struct {
	Key     string
	Message string
	Errors  *[]Error
	Status  *int16
}

func (e *Error) Error() string {
	return e.Message
}

func NewCustom(key string, message string, errors *[]Error, status *int16) *Error {
	return &Error{
		Key:     key,
		Message: message,
		Errors:  errors,
		Status:  status,
	}
}
