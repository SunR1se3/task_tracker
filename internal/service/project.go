package service

import (
	"github.com/google/uuid"
	"strings"
	"task_tracker/internal/domain"
	"task_tracker/internal/repository"
)

type ProjectService struct {
	repo repository.Project
}

func NewProjectService(r repository.Project) *ProjectService {
	return &ProjectService{repo: r}
}

func (s *ProjectService) CreateProject(formData *domain.ProjectCreateForm, userId uuid.UUID) (*uuid.UUID, error) {
	data := new(domain.Project)
	err := formData.Prepare(data)
	if err != nil {
		return nil, err
	}
	err = s.repo.CreateProject(data, userId)
	if err != nil {
		return nil, err
	}
	return &data.Id, nil
}

func (s *ProjectService) GetProjectById(id uuid.UUID) (*domain.Project, error) {
	return s.repo.GetProjectById(id)
}

func (s *ProjectService) GetProjectsUserId(userId uuid.UUID) ([]domain.Project, error) {
	return s.repo.GetProjectsUserId(userId)
}

func (s *ProjectService) AddUserToTeam(formData *domain.AddUserToTeamForm) error {
	return s.repo.AddUserToTeam(*formData.UserId, *formData.ProjectId)
}

func (s *ProjectService) SetUserProjectRole(formData *domain.AddUserToTeamForm) error {
	return s.repo.SetUserProjectRole(formData.UserId, formData.ProjectId, formData.ProjectRoleId)
}

func (s *ProjectService) GetProjectTeam(projectId uuid.UUID) ([]domain.Teammate, error) {
	team, err := s.repo.GetProjectTeam(projectId)
	if err != nil {
		return nil, err
	}
	for i, teammate := range team {
		projectRole, err := s.repo.GetProjectRoleForUser(projectId, teammate.Id)
		if err != nil {
			if !strings.Contains(err.Error(), "no rows") {
				return nil, err
			}
			continue
		}
		team[i].ProjectRole = &projectRole
	}
	return team, nil
}

func (s *ProjectService) GetProjectRoles() []domain.ProjectRole {
	return s.repo.GetProjectRoles()
}
