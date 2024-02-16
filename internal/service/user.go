package service

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"task_tracker/internal/constants"
	"task_tracker/internal/domain"
	"task_tracker/internal/errors"
	"task_tracker/internal/helper"
	"task_tracker/internal/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(r repository.User) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) CreateUser(formData *domain.UserCreateForm) (*uuid.UUID, error) {
	data := new(domain.User)
	formData.Prepare(data)
	exists, _ := s.repo.AlreadyExists(formData.Login)
	if exists {
		return nil, errors.UserAlreadyExists()
	}
	err := s.repo.CreateUser(data)
	if err != nil {
		return nil, err
	}
	return &data.Id, nil
}

func (s *UserService) GetUserById(id uuid.UUID) (*domain.User, error) {
	return s.repo.GetUserById(id)
}

func (s *UserService) GetUserDTOById(id uuid.UUID) (*domain.UserDTO, error) {
	positions, err := Services.Position.GetUserPositions(id)
	if err != nil {
		return nil, err
	}
	departments, err := Services.Department.GetUserDepartments(id)
	if err != nil {
		return nil, err
	}
	specializations, err := Services.Specialization.GetUserSpecializations(id)
	if err != nil {
		return nil, err
	}
	user, err := s.repo.GetUserDTOById(id)
	if err != nil {
		return nil, err
	}
	user.Positions = positions
	user.Departments = departments
	user.Specializations = specializations
	return user, err
}

func (s *UserService) GetUsersDTO() ([]domain.UserDTO, error) {
	users, err := s.repo.GetUsersDTO()
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(users); i++ {
		positions, err := Services.Position.GetUserPositions(users[i].Id)
		if err != nil {
			return nil, err
		}
		departments, err := Services.Department.GetUserDepartments(users[i].Id)
		if err != nil {
			return nil, err
		}
		specializations, err := Services.Specialization.GetUserSpecializations(users[i].Id)
		if err != nil {
			return nil, err
		}
		users[i].Positions = positions
		users[i].Departments = departments
		users[i].Specializations = specializations
	}
	return users, err
}

func (s *UserService) AdminUsersTable() (*string, error) {
	users, err := s.GetUsersDTO()
	if err != nil {
		return nil, err
	}
	departments, err := Services.Department.GetAll()
	if err != nil {
		return nil, err
	}
	specializations, err := Services.Specialization.GetAll()
	if err != nil {
		return nil, err
	}
	positions, err := Services.Position.GetAll()
	if err != nil {
		return nil, err
	}

	return helper.HtmlRenderProcess("./views/admin_pages/users/table.html", "table", map[string]interface{}{
		"users":           users,
		"departments":     departments,
		"specializations": specializations,
		"positions":       positions,
	})
}

func (s *UserService) ChangePassword(formData *domain.ChangePasswordForm, userId *uuid.UUID) []error {
	var errs []error
	user, err := s.repo.GetUserById(*userId)
	if err != nil {
		errs = append(errs, err)
		return errs
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(formData.OldPassword))
	if err != nil {
		errs = append(errs, errors.InvalidPassword(helper.GetJsonTag("OldPassword", *formData)))
		return errs
	}
	errs = formData.Prepare()
	if len(errs) > 0 {
		return errs
	}
	err = s.repo.ChangePassword(formData.NewPassword, userId)
	if err != nil {
		errs = append(errs, err)
	}
	return errs
}

func (s *UserService) GetEditUserModalForm(id uuid.UUID) (*string, error) {
	user, err := s.GetUserDTOById(id)
	if err != nil {
		return nil, err
	}
	departments, err := Services.Department.GetAll()
	if err != nil {
		return nil, err
	}
	specializations, err := Services.Specialization.GetAll()
	if err != nil {
		return nil, err
	}
	positions, err := Services.Position.GetAll()
	if err != nil {
		return nil, err
	}
	return helper.HtmlRenderProcess("./views/admin_pages/users/modal_edit.html", "modal_edit", map[string]interface{}{
		"user":            user,
		"departments":     departments,
		"specializations": specializations,
		"positions":       positions,
		"systemRoles":     constants.SystemRoles,
	})
}

func (s *UserService) EditUser(id uuid.UUID, formData *domain.UserEditForm) error {
	data, err := s.GetUserById(id)
	if err != nil {
		return err
	}
	formData.Prepare(data)
	err = s.repo.EditUser(data)
	return err
}

func (s *UserService) DisableUser(userId uuid.UUID, disable bool) error {
	return s.repo.DisableUser(userId, disable)
}

func (s *UserService) UserPicker() ([]domain.UserPicker, error) {
	return s.repo.UserPicker()
}
