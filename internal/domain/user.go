package domain

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Id                 uuid.UUID   `json:"id" db:"id"`
	Login              string      `json:"login" db:"login"`
	Password           string      `json:"password" db:"password"`
	Firstname          string      `json:"firstname" db:"firstname"`
	Middlename         string      `json:"middlename" db:"middlename"`
	Lastname           string      `json:"lastname" db:"lastname"`
	IsActive           bool        `json:"isActive" db:"is_active"`
	AccountDisableTime *time.Time  `json:"accountDisableTime" db:"account_disable_time"`
	SystemRole         int         `json:"systemRole" db:"system_role"`
	Positions          []uuid.UUID `json:"positions" db:"positions"`
	Departments        []uuid.UUID `json:"departments" db:"departments"`
	Specializations    []uuid.UUID `json:"specializations" db:"specializations"`
	CreatedAt          time.Time   `json:"createdAt" db:"created_at"`
	UpdatedAt          *time.Time  `json:"updatedAt" db:"updated_at"`
}

type UserDTO struct {
	Id                 uuid.UUID        `json:"id" db:"id"`
	Login              string           `json:"login" db:"login"`
	Firstname          string           `json:"firstname" db:"firstname"`
	Middlename         string           `json:"middlename" db:"middlename"`
	Lastname           string           `json:"lastname" db:"lastname"`
	IsActive           bool             `json:"isActive" db:"is_active"`
	AccountDisableTime *time.Time       `json:"accountDisableTime" db:"account_disable_time"`
	SystemRole         string           `json:"systemRole" db:"system_role"`
	Positions          []Position       `json:"positions" db:"positions"`
	Departments        []Department     `json:"departments" db:"departments"`
	Specializations    []Specialization `json:"specializations" db:"specializations"`
	CreatedAt          time.Time        `json:"createdAt" db:"created_at"`
}

type UserCreateForm struct {
	Login           string      `json:"login" db:"login"`
	Password        string      `json:"password" db:"password"`
	Firstname       string      `json:"firstname" db:"firstname"`
	Middlename      string      `json:"middlename" db:"middlename"`
	Lastname        string      `json:"lastname" db:"lastname"`
	Positions       []uuid.UUID `json:"positions" db:"positions"`
	Departments     []uuid.UUID `json:"departments" db:"departments"`
	Specializations []uuid.UUID `json:"specializations" db:"specializations"`
}

func (f *UserCreateForm) Prepare(m *User) {
	m.Id = uuid.New()
	m.Login = f.Login

	bytePassword, _ := bcrypt.GenerateFromPassword([]byte(f.Password), bcrypt.DefaultCost)
	hashPassword := string(bytePassword[:])
	m.Password = hashPassword

	m.Firstname = f.Firstname
	m.Middlename = f.Middlename
	m.Lastname = f.Lastname
	m.IsActive = true
	m.SystemRole = 1
	m.Positions = f.Positions
	m.Departments = f.Departments
	m.Specializations = f.Specializations
	m.CreatedAt = time.Now()
}
