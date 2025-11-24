package postgres

import (
	"context"

	"github.com/Swarnadip-Dey/Collaborative-taskmanager/internal/models"
	"github.com/Swarnadip-Dey/Collaborative-taskmanager/internal/repository"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *userRepository) GetByID(ctx context.Context, id uint) (*models.User, error) {
	var user models.User
	if err := r.db.WithContext(ctx).First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

type workspaceRepository struct {
	db *gorm.DB
}

func NewWorkspaceRepository(db *gorm.DB) repository.WorkspaceRepository {
	return &workspaceRepository{db: db}
}

func (r *workspaceRepository) Create(ctx context.Context, workspace *models.Workspace) error {
	return r.db.WithContext(ctx).Create(workspace).Error
}

func (r *workspaceRepository) GetByID(ctx context.Context, id uint) (*models.Workspace, error) {
	var workspace models.Workspace
	if err := r.db.WithContext(ctx).Preload("Owner").First(&workspace, id).Error; err != nil {
		return nil, err
	}
	return &workspace, nil
}

func (r *workspaceRepository) ListByUserID(ctx context.Context, userID uint) ([]models.Workspace, error) {
	var workspaces []models.Workspace
	// Assuming logic: workspaces owned by user.
	// In a real app, might also include workspaces where user is a member.
	if err := r.db.WithContext(ctx).Where("owner_id = ?", userID).Find(&workspaces).Error; err != nil {
		return nil, err
	}
	return workspaces, nil
}

type projectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) repository.ProjectRepository {
	return &projectRepository{db: db}
}

func (r *projectRepository) Create(ctx context.Context, project *models.Project) error {
	return r.db.WithContext(ctx).Create(project).Error
}

func (r *projectRepository) GetByID(ctx context.Context, id uint) (*models.Project, error) {
	var project models.Project
	if err := r.db.WithContext(ctx).Preload("Workspace").First(&project, id).Error; err != nil {
		return nil, err
	}
	return &project, nil
}

func (r *projectRepository) ListByWorkspaceID(ctx context.Context, workspaceID uint) ([]models.Project, error) {
	var projects []models.Project
	if err := r.db.WithContext(ctx).Where("workspace_id = ?", workspaceID).Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) repository.TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) Create(ctx context.Context, task *models.Task) error {
	return r.db.WithContext(ctx).Create(task).Error
}

func (r *taskRepository) GetByID(ctx context.Context, id uint) (*models.Task, error) {
	var task models.Task
	if err := r.db.WithContext(ctx).Preload("Assignee").Preload("Project").First(&task, id).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *taskRepository) Update(ctx context.Context, task *models.Task) error {
	return r.db.WithContext(ctx).Save(task).Error
}

func (r *taskRepository) ListByProjectID(ctx context.Context, projectID uint) ([]models.Task, error) {
	var tasks []models.Task
	if err := r.db.WithContext(ctx).Where("project_id = ?", projectID).Preload("Assignee").Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

type taskHistoryRepository struct {
	db *gorm.DB
}

func NewTaskHistoryRepository(db *gorm.DB) repository.TaskHistoryRepository {
	return &taskHistoryRepository{db: db}
}

func (r *taskHistoryRepository) Create(ctx context.Context, history *models.TaskHistory) error {
	return r.db.WithContext(ctx).Create(history).Error
}

func (r *taskHistoryRepository) ListByTaskID(ctx context.Context, taskID uint) ([]models.TaskHistory, error) {
	var history []models.TaskHistory
	if err := r.db.WithContext(ctx).Where("task_id = ?", taskID).Order("created_at desc").Find(&history).Error; err != nil {
		return nil, err
	}
	return history, nil
}

type Repository struct {
	db          *gorm.DB
	users       repository.UserRepository
	workspaces  repository.WorkspaceRepository
	projects    repository.ProjectRepository
	tasks       repository.TaskRepository
	taskHistory repository.TaskHistoryRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db:          db,
		users:       NewUserRepository(db),
		workspaces:  NewWorkspaceRepository(db),
		projects:    NewProjectRepository(db),
		tasks:       NewTaskRepository(db),
		taskHistory: NewTaskHistoryRepository(db),
	}
}

func (r *Repository) Users() repository.UserRepository {
	return r.users
}

func (r *Repository) Workspaces() repository.WorkspaceRepository {
	return r.workspaces
}

func (r *Repository) Projects() repository.ProjectRepository {
	return r.projects
}

func (r *Repository) Tasks() repository.TaskRepository {
	return r.tasks
}

func (r *Repository) TaskHistory() repository.TaskHistoryRepository {
	return r.taskHistory
}
