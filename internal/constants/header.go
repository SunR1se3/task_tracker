package constants

var (
	Headers = map[int]interface{}{
		0: AdminHeader,
		1: DefaultHeader,
	}
	DefaultHeader = map[string]string{
		"team":     "Команда",
		"boards":   "Доски",
		"projects": "Проекты",
		"reports":  "Отчеты",
	}
	AdminHeader = map[string]string{
		"projects":     "Проекты",
		"/admin/users": "Пользователи",
	}
)
