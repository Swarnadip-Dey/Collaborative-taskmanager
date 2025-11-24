package repository

import (
	"context"

	"github.com/Swarnadip-Dey/Collaborative-taskmanager/internal/models"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	GetByID(ctx context.Context, id uint) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
}

type WorkspaceRepository interface {
	Create(ctx context.Context, workspace *models.Workspace) error
	GetByID(ctx context.Context, id uint) (*models.Workspace, error)
	ListByUserID(ctx context.Context, userID uint) ([]models.Workspace, error)
}

type ProjectRepository interface {
	Create(ctx context.Context, project *models.Project) error
	GetByID(ctx context.Context, id uint) (*models.Project, error)
	ListByWorkspaceID(ctx context.Context, workspaceID uint) ([]models.Project, error)
}

type TaskRepository interface {
	Create(ctx context.Context, task *models.Task) error
	GetByID(ctx context.Context, id uint) (*models.Task, error)
	Update(ctx context.Context, task *models.Task) error
	ListByProjectID(ctx context.Context, projectID uint) ([]models.Task, error)
}

type TaskHistoryRepository interface {
	Create(ctx context.Context, history *models.TaskHistory) error
	ListByTaskID(ctx context.Context, taskID uint) ([]models.TaskHistory, error)
}

type Repository interface {
	Users() UserRepository
	Workspaces() WorkspaceRepository
	Projects() ProjectRepository
	Tasks() TaskRepository
	TaskHistory() TaskHistoryRepository
}
