# API Documentation - Role-Based Endpoints

## Overview
The Collaborative Task Manager API now has three role-based access levels with dedicated controllers and services.

## Role Hierarchy
1. **Admin** - Full access to all endpoints
2. **Manager** - Can manage workspaces, projects, and assign tasks
3. **Developer** - Can create and update tasks, view projects

---

## Authentication Endpoints (Public)

### POST /api/register
Register a new user
- **Body**: `{ "username": "string", "email": "string", "password": "string", "role": "admin|manager|dev" }`
- **Response**: JWT token + user object

### POST /api/login
Login user
- **Body**: `{ "email": "string", "password": "string" }`
- **Response**: JWT token + user object

### GET /api/profile
Get current user profile (requires authentication)
- **Headers**: `Authorization: Bearer <token>`
- **Response**: User object

---

## Manager Endpoints (Requires Manager or Admin Role)

### POST /api/manager/workspaces
Create a new workspace
- **Headers**: `Authorization: Bearer <token>`
- **Body**: `{ "name": "string" }`
- **Response**: Workspace object

### POST /api/manager/projects
Create a new project
- **Headers**: `Authorization: Bearer <token>`
- **Body**: `{ "name": "string", "workspace_id": number }`
- **Response**: Project object

### GET /api/manager/workspaces/:workspace_id/projects
List all projects in a workspace
- **Headers**: `Authorization: Bearer <token>`
- **Response**: Array of projects

### PUT /api/manager/tasks/:id/assign
Assign a task to a developer
- **Headers**: `Authorization: Bearer <token>`
- **Body**: `{ "assignee_id": number }`
- **Response**: Updated task object

---

## Developer Endpoints (Requires Authentication)

### POST /api/dev/tasks
Create a new task (auto-assigned to creator)
- **Headers**: `Authorization: Bearer <token>`
- **Body**: 
```json
{
  "title": "string",
  "description": "string",
  "status": "TODO|IN_PROGRESS|DONE",
  "priority": "LOW|MEDIUM|HIGH",
  "project_id": number
}
```
- **Response**: Task object

### GET /api/dev/tasks/:id
Get task by ID
- **Headers**: `Authorization: Bearer <token>`
- **Response**: Task object

### PUT /api/dev/tasks/:id
Update a task
- **Headers**: `Authorization: Bearer <token>`
- **Body**: 
```json
{
  "title": "string",
  "description": "string",
  "status": "TODO|IN_PROGRESS|DONE",
  "priority": "LOW|MEDIUM|HIGH"
}
```
- **Response**: Updated task object

### GET /api/dev/projects/:project_id/tasks
List all tasks in a project
- **Headers**: `Authorization: Bearer <token>`
- **Response**: Array of tasks

### GET /api/dev/projects/:id
Get project by ID
- **Headers**: `Authorization: Bearer <token>`
- **Response**: Project object

---

## Admin Endpoints (Requires Admin Role)

### GET /api/admin/users
List all users (placeholder)
- **Headers**: `Authorization: Bearer <token>`
- **Response**: Admin endpoint message

---

## Services Layer

### WorkspaceService
- `CreateWorkspace(ctx, name, ownerID)` - Create workspace
- `GetWorkspace(ctx, id)` - Get workspace by ID
- `ListUserWorkspaces(ctx, userID)` - List user's workspaces

### ProjectService
- `CreateProject(ctx, name, workspaceID)` - Create project
- `GetProject(ctx, id)` - Get project by ID
- `ListWorkspaceProjects(ctx, workspaceID)` - List workspace projects

### TaskService
- `CreateTask(ctx, input)` - Create task
- `GetTask(ctx, id)` - Get task by ID
- `UpdateTask(ctx, id, input)` - Update task
- `ListProjectTasks(ctx, projectID)` - List project tasks
- `AssignTask(ctx, taskID, assigneeID)` - Assign task to user

---

## Swagger Documentation

Access the interactive API documentation at:
```
http://localhost:8080/swagger/index.html
```

All endpoints are documented with request/response schemas and can be tested directly from the Swagger UI.

---

## Example Workflow

1. **Register as Manager**:
   ```bash
   curl -X POST http://localhost:8080/api/register \
     -H "Content-Type: application/json" \
     -d '{"username":"manager1","email":"manager@example.com","password":"password123","role":"manager"}'
   ```

2. **Login**:
   ```bash
   curl -X POST http://localhost:8080/api/login \
     -H "Content-Type: application/json" \
     -d '{"email":"manager@example.com","password":"password123"}'
   ```

3. **Create Workspace** (use token from login):
   ```bash
   curl -X POST http://localhost:8080/api/manager/workspaces \
     -H "Authorization: Bearer <token>" \
     -H "Content-Type: application/json" \
     -d '{"name":"My Team"}'
   ```

4. **Create Project**:
   ```bash
   curl -X POST http://localhost:8080/api/manager/projects \
     -H "Authorization: Bearer <token>" \
     -H "Content-Type: application/json" \
     -d '{"name":"Project Alpha","workspace_id":1}'
   ```

5. **Developer Creates Task**:
   ```bash
   curl -X POST http://localhost:8080/api/dev/tasks \
     -H "Authorization: Bearer <dev-token>" \
     -H "Content-Type: application/json" \
     -d '{"title":"Fix bug","description":"Fix login issue","project_id":1}'
   ```
