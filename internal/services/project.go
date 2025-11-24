package services

import (
	"context"
	"fmt"

	"github.com/Swarnadip-Dey/Collaborative-taskmanager/internal/models"
	"github.com/Swarnadip-Dey/Collaborative-taskmanager/internal/repository"
)

type ProjectService struct {
	repo repository.Repository
}

func NewProjectService(repo repository.Repository) *ProjectService {
	return &ProjectService{repo: repo}
}

func (s *ProjectService) CreateProject(ctx context.Context, name string, workspaceID uint) (*models.Project, error) {
	project := &models.Project{
		Name:        name,
		WorkspaceID: workspaceID,
	}

	if err := s.repo.Projects().Create(ctx, project); err != nil {
		return nil, fmt.Errorf("failed to create project: %w", err)
	}

	return project, nil
}

func (s *ProjectService) GetProject(ctx context.Context, id uint) (*models.Project, error) {
	return s.repo.Projects().GetByID(ctx, id)
}

func (s *ProjectService) ListWorkspaceProjects(ctx context.Context, workspaceID uint) ([]models.Project, error) {
	return s.repo.Projects().ListByWorkspaceID(ctx, workspaceID)
}
