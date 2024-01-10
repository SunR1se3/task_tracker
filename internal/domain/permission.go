package domain

type Permission struct {
	EntityBase
	Title string `json:"title" db:"title"`
	Code  string `json:"code" db:"code"`
}
