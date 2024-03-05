package errors

import "fmt"

func EndDateSprintError(field string) error {
	return New(field, fmt.Sprintf("Дату окончания спринта нельзя поставить раньше начала спринта"))
}
