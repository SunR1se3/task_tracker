package domain

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"task_tracker/internal/errors"
	"task_tracker/internal/helper"
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

type UserEditForm struct {
	Login           string      `json:"login" db:"login"`
	Firstname       string      `json:"firstname" db:"firstname"`
	Middlename      string      `json:"middlename" db:"middlename"`
	Lastname        string      `json:"lastname" db:"lastname"`
	SystemRole      int         `json:"systemRole" db:"system_role"`
	Positions       []uuid.UUID `json:"positions" db:"positions"`
	Departments     []uuid.UUID `json:"departments" db:"departments"`
	Specializations []uuid.UUID `json:"specializations" db:"specializations"`
}

type ChangePasswordForm struct {
	OldPassword         string `json:"oldPassword"`
	NewPassword         string `json:"newPassword"`
	RepeatedNewPassword string `json:"repeatedNewPassword"`
}

func (f *ChangePasswordForm) Prepare() []error {
	errs := f.Validate()
	if len(errs) > 0 {
		return errs
	}
	bytePassword, _ := bcrypt.GenerateFromPassword([]byte(f.OldPassword), bcrypt.DefaultCost)
	hashPassword := string(bytePassword[:])
	f.OldPassword = hashPassword

	bytePassword, _ = bcrypt.GenerateFromPassword([]byte(f.NewPassword), bcrypt.DefaultCost)
	hashPassword = string(bytePassword[:])
	f.NewPassword = hashPassword

	bytePassword, _ = bcrypt.GenerateFromPassword([]byte(f.RepeatedNewPassword), bcrypt.DefaultCost)
	hashPassword = string(bytePassword[:])
	f.RepeatedNewPassword = hashPassword
	return errs
}

func (f *ChangePasswordForm) Validate() []error {
	errs := []error{}
	if f.OldPassword == "" {
		errs = append(errs, errors.RequiredFiledError(helper.GetJsonTag("OldPassword", *f)))
	}
	if f.NewPassword == "" {
		errs = append(errs, errors.RequiredFiledError(helper.GetJsonTag("NewPassword", *f)))
	}
	if f.RepeatedNewPassword == "" {
		errs = append(errs, errors.RequiredFiledError(helper.GetJsonTag("RepeatedNewPassword", *f)))
	}
	if f.NewPassword != f.RepeatedNewPassword {
		errs = append(errs, errors.NotEqualPassword(helper.GetJsonTag("NewPassword", *f)))
		errs = append(errs, errors.NotEqualPassword(helper.GetJsonTag("RepeatedNewPassword", *f)))
	}
	return errs
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

func (f *UserEditForm) Prepare(m *User) {
	m.Login = f.Login
	m.Firstname = f.Firstname
	m.Middlename = f.Middlename
	m.Lastname = f.Lastname
	m.SystemRole = f.SystemRole
	m.Positions = f.Positions
	m.Departments = f.Departments
	m.Specializations = f.Specializations
	t := time.Now()
	m.UpdatedAt = &t
}
