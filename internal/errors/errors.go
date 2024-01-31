package errors

type Error struct {
	Code string `json:"code"`
	Msg  string `json:"msg" xml:"msg" form:"msg"`
}

func (e Error) Error() string {
	return e.Code + "|" + e.Msg
}

func New(code, msg string) error {
	return &Error{Code: code, Msg: msg}
}

func RequiredFiledError(field string) error {
	return New(field, "Поле не может быть пустым")
}
