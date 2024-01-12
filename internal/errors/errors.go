package errors

type Error struct {
	Msg string `json:"msg" xml:"msg" form:"msg"`
}

func (e Error) Error() string {
	return e.Msg
}

func New(msg string) error {
	return &Error{Msg: msg}
}
