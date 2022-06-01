package goerr

type Error struct {
	Key     string     `json:"key"`
	Message string     `json:"message"`
	Errors  []SubError `json:"errors"`
	Status  int        `json:"status"`
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) InsertSubError(errors ...SubError) {
	e.Errors = append(e.Errors, errors...)
}
