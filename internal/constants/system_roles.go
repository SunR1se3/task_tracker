package constants

const (
	ROLE_ADMIN = 0
	ROLE_USER  = 1
)

var (
	SystemRoles = map[int]string{
		ROLE_ADMIN: "Админ",
		ROLE_USER:  "Пользователь",
	}
)
