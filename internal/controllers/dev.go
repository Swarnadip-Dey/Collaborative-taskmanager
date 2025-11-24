package controllers

import (
	"net/http"
	"strconv"

	"github.com/Swarnadip-Dey/Collaborative-taskmanager/internal/models"
	"github.com/Swarnadip-Dey/Collaborative-taskmanager/internal/services"
	"github.com/gin-gonic/gin"
)

type DevController struct {
	taskService    *services.TaskService
	projectService *services.ProjectService
}

func NewDevController(
	taskService *services.TaskService,
	projectService *services.ProjectService,
) *DevController {
	return &DevController{
		taskService:    taskService,
		projectService: projectService,
	}
}

type CreateTaskRequest struct {
	Title       string              `json:"title" binding:"required"`
	Description string              `json:"description"`
	Status      models.TaskStatus   `json:"status"`
	Priority    models.TaskPriority `json:"priority"`
	ProjectID   uint                `json:"project_id" binding:"required"`
}

type UpdateTaskRequest struct {
	Title       string              `json:"title"`
	Description string              `json:"description"`
	Status      models.TaskStatus   `json:"status"`
	Priority    models.TaskPriority `json:"priority"`
}

// CreateTask godoc
// @Summary Create a new task
// @Description Developer can create a new task in a project
// @Tags developer
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body CreateTaskRequest true "Task details"
// @Success 201 {object} models.Task
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/dev/tasks [post]
func (dc *DevController) CreateTask(c *gin.Context) {
	var req CreateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("user_id")
	assigneeID := userID.(uint)

	input := services.CreateTaskInput{
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		Priority:    req.Priority,
		AssigneeID:  &assigneeID,
		ProjectID:   req.ProjectID,
	}

	task, err := dc.taskService.CreateTask(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, task)
}

// GetTask godoc
// @Summary Get task by ID
// @Description Developer can view a task by its ID
// @Tags developer
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Task ID"
// @Success 200 {object} models.Task
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/dev/tasks/{id} [get]
func (dc *DevController) GetTask(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
		return
	}

	task, err := dc.taskService.GetTask(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

// UpdateTask godoc
// @Summary Update a task
// @Description Developer can update task details (title, description, status, priority)
// @Tags developer
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Task ID"
// @Param request body UpdateTaskRequest true "Updated task details"
// @Success 200 {object} models.Task
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/dev/tasks/{id} [put]
func (dc *DevController) UpdateTask(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
		return
	}

	var req UpdateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input := services.UpdateTaskInput{}
	if req.Title != "" {
		input.Title = &req.Title
	}
	if req.Description != "" {
		input.Description = &req.Description
	}
	if req.Status != "" {
		input.Status = &req.Status
	}
	if req.Priority != "" {
		input.Priority = &req.Priority
	}

	task, err := dc.taskService.UpdateTask(c.Request.Context(), uint(id), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

// ListProjectTasks godoc
// @Summary List tasks in a project
// @Description Developer can view all tasks in a specific project
// @Tags developer
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Project ID"
// @Success 200 {array} models.Task
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/dev/projects/{id}/tasks [get]
func (dc *DevController) ListProjectTasks(c *gin.Context) {
	projectID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid project ID"})
		return
	}

	tasks, err := dc.taskService.ListProjectTasks(c.Request.Context(), uint(projectID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// GetProject godoc
// @Summary Get project by ID
// @Description Developer can view a project by its ID
// @Tags developer
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Project ID"
// @Success 200 {object} models.Project
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/dev/projects/{id} [get]
func (dc *DevController) GetProject(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid project ID"})
		return
	}

	project, err := dc.projectService.GetProject(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "project not found"})
		return
	}

	c.JSON(http.StatusOK, project)
}
