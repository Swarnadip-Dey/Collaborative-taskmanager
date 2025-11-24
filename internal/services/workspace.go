package services

import (
	"context"
	"fmt"

	"github.com/Swarnadip-Dey/Collaborative-taskmanager/internal/models"
	"github.com/Swarnadip-Dey/Collaborative-taskmanager/internal/repository"
)

type WorkspaceService struct {
	repo repository.Repository
}

func NewWorkspaceService(repo repository.Repository) *WorkspaceService {
	return &WorkspaceService{repo: repo}
}

func (s *WorkspaceService) CreateWorkspace(ctx context.Context, name string, ownerID uint) (*models.Workspace, error) {
	workspace := &models.Workspace{
		Name:    name,
		OwnerID: ownerID,
	}

	if err := s.repo.Workspaces().Create(ctx, workspace); err != nil {
		return nil, fmt.Errorf("failed to create workspace: %w", err)
	}

	return workspace, nil
}

func (s *WorkspaceService) GetWorkspace(ctx context.Context, id uint) (*models.Workspace, error) {
	return s.repo.Workspaces().GetByID(ctx, id)
}

func (s *WorkspaceService) ListUserWorkspaces(ctx context.Context, userID uint) ([]models.Workspace, error) {
	return s.repo.Workspaces().ListByUserID(ctx, userID)
}
