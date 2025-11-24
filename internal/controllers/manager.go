package controllers

import (
	"net/http"
	"strconv"

	"github.com/Swarnadip-Dey/Collaborative-taskmanager/internal/services"
	"github.com/gin-gonic/gin"
)

type ManagerController struct {
	workspaceService *services.WorkspaceService
	projectService   *services.ProjectService
	taskService      *services.TaskService
}

func NewManagerController(
	workspaceService *services.WorkspaceService,
	projectService *services.ProjectService,
	taskService *services.TaskService,
) *ManagerController {
	return &ManagerController{
		workspaceService: workspaceService,
		projectService:   projectService,
		taskService:      taskService,
	}
}

type CreateWorkspaceRequest struct {
	Name string `json:"name" binding:"required"`
}

type CreateProjectRequest struct {
	Name        string `json:"name" binding:"required"`
	WorkspaceID uint   `json:"workspace_id" binding:"required"`
}

type AssignTaskRequest struct {
	AssigneeID uint `json:"assignee_id" binding:"required"`
}

// CreateWorkspace godoc
// @Summary Create a new workspace
// @Description Manager/Admin can create a new workspace (team)
// @Tags manager
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body CreateWorkspaceRequest true "Workspace details"
// @Success 201 {object} models.Workspace
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/manager/workspaces [post]
func (mc *ManagerController) CreateWorkspace(c *gin.Context) {
	var req CreateWorkspaceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("user_id")

	workspace, err := mc.workspaceService.CreateWorkspace(c.Request.Context(), req.Name, userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, workspace)
}

// CreateProject godoc
// @Summary Create a new project
// @Description Manager/Admin can create a new project in a workspace
// @Tags manager
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body CreateProjectRequest true "Project details"
// @Success 201 {object} models.Project
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/manager/projects [post]
func (mc *ManagerController) CreateProject(c *gin.Context) {
	var req CreateProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	project, err := mc.projectService.CreateProject(c.Request.Context(), req.Name, req.WorkspaceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, project)
}

// AssignTask godoc
// @Summary Assign a task to a user
// @Description Manager/Admin can assign a task to a developer
// @Tags manager
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Task ID"
// @Param request body AssignTaskRequest true "Assignee details"
// @Success 200 {object} models.Task
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/manager/tasks/{id}/assign [put]
func (mc *ManagerController) AssignTask(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
		return
	}

	var req AssignTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := mc.taskService.AssignTask(c.Request.Context(), uint(id), req.AssigneeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

// ListWorkspaceProjects godoc
// @Summary List projects in a workspace
// @Description Manager/Admin can view all projects in a workspace
// @Tags manager
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param workspace_id path int true "Workspace ID"
// @Success 200 {array} models.Project
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/manager/workspaces/{workspace_id}/projects [get]
func (mc *ManagerController) ListWorkspaceProjects(c *gin.Context) {
	workspaceID, err := strconv.ParseUint(c.Param("workspace_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid workspace ID"})
		return
	}

	projects, err := mc.projectService.ListWorkspaceProjects(c.Request.Context(), uint(workspaceID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, projects)
}
