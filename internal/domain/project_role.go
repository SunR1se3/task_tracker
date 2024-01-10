package domain

type ProjectRole struct {
	EntityBase
	Title       string       `json:"title" db:"title"`
	Code        string       `json:"code" db:"code"`
	Permissions []Permission `json:"permissions" db:"permissions"`
}
