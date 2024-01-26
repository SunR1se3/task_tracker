package errors

type ErrorHandlerInterface interface {
	Add(err ...error)
	Get() []string
}
type ErrorHandler struct {
	ErrList []string
}

func (e *ErrorHandler) Add(err ...error) {
	for _, r := range err {
		e.ErrList = append(e.ErrList, r.Error())
	}
}

func (e *ErrorHandler) Get() []string {
	return e.ErrList
}
