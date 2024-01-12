package errors

func UserAlreadyExists() error {
	return New("Пользователь с таким логином существует")
}
