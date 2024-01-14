package domain

type AuthForm struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
