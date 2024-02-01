package service

import (
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"task_tracker/internal/domain"
	"task_tracker/internal/repository"
	"time"
)

type AuthService struct {
	repo repository.User
}

func NewAuthService(r repository.User) *AuthService {
	return &AuthService{repo: r}
}

func (s *AuthService) Auth(formData *domain.AuthForm) (*jwt.Token, error) {
	user, err := s.repo.GetUserByLogin(formData.Login)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(formData.Password))
	if err != nil {
		return nil, err
	}
	claims := jwt.MapClaims{
		"id":         user.Id,
		"login":      user.Login,
		"firstname":  user.Firstname,
		"lastname":   user.Lastname,
		"middlename": user.Middlename,
		"role":       user.SystemRole,
		"exp":        time.Now().Add(time.Hour * 72).Unix(),
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims), err

}
