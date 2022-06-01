package goerr

type SubError struct {
	Key     string `json:"key"`
	Message string `json:"message"`
}

func (e *Error) SubError() string {
	return e.Message
}
