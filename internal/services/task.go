package services

import (
	"context"
	"fmt"

	"github.com/Swarnadip-Dey/Collaborative-taskmanager/internal/models"
	"github.com/Swarnadip-Dey/Collaborative-taskmanager/internal/repository"
)

type TaskService struct {
	repo repository.Repository
}

func NewTaskService(repo repository.Repository) *TaskService {
	return &TaskService{repo: repo}
}

type CreateTaskInput struct {
	Title       string
	Description string
	Status      models.TaskStatus
	Priority    models.TaskPriority
	AssigneeID  *uint
	ProjectID   uint
}

type UpdateTaskInput struct {
	Title       *string
	Description *string
	Status      *models.TaskStatus
	Priority    *models.TaskPriority
	AssigneeID  *uint
}

func (s *TaskService) CreateTask(ctx context.Context, input CreateTaskInput) (*models.Task, error) {
	// Set defaults
	if input.Status == "" {
		input.Status = models.TaskStatusTodo
	}
	if input.Priority == "" {
		input.Priority = models.TaskPriorityMedium
	}

	task := &models.Task{
		Title:       input.Title,
		Description: input.Description,
		Status:      input.Status,
		Priority:    input.Priority,
		AssigneeID:  input.AssigneeID,
		ProjectID:   input.ProjectID,
	}

	if err := s.repo.Tasks().Create(ctx, task); err != nil {
		return nil, fmt.Errorf("failed to create task: %w", err)
	}

	return task, nil
}

func (s *TaskService) GetTask(ctx context.Context, id uint) (*models.Task, error) {
	return s.repo.Tasks().GetByID(ctx, id)
}

func (s *TaskService) UpdateTask(ctx context.Context, id uint, input UpdateTaskInput) (*models.Task, error) {
	task, err := s.repo.Tasks().GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("task not found: %w", err)
	}

	// Update fields if provided
	if input.Title != nil {
		task.Title = *input.Title
	}
	if input.Description != nil {
		task.Description = *input.Description
	}
	if input.Status != nil {
		task.Status = *input.Status
	}
	if input.Priority != nil {
		task.Priority = *input.Priority
	}
	if input.AssigneeID != nil {
		task.AssigneeID = input.AssigneeID
	}

	if err := s.repo.Tasks().Update(ctx, task); err != nil {
		return nil, fmt.Errorf("failed to update task: %w", err)
	}

	return task, nil
}

func (s *TaskService) ListProjectTasks(ctx context.Context, projectID uint) ([]models.Task, error) {
	return s.repo.Tasks().ListByProjectID(ctx, projectID)
}

func (s *TaskService) AssignTask(ctx context.Context, taskID uint, assigneeID uint) (*models.Task, error) {
	task, err := s.repo.Tasks().GetByID(ctx, taskID)
	if err != nil {
		return nil, fmt.Errorf("task not found: %w", err)
	}

	task.AssigneeID = &assigneeID

	if err := s.repo.Tasks().Update(ctx, task); err != nil {
		return nil, fmt.Errorf("failed to assign task: %w", err)
	}

	return task, nil
}
