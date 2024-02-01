package errors

func UserAlreadyExists() error {
	return New("unique-login", "Пользователь с таким логином существует")
}

func NotEqualPassword(field string) error {
	return New(field, "Пароли должны совпадать")
}

func InvalidPassword(field string) error {
	return New(field, "Неверный логин или пароль")
}
