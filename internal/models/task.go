package models

import "time"

type TaskStatus string
type TaskPriority string

const (
	TaskStatusTodo       TaskStatus = "TODO"
	TaskStatusInProgress TaskStatus = "IN_PROGRESS"
	TaskStatusDone       TaskStatus = "DONE"
)

const (
	TaskPriorityLow    TaskPriority = "LOW"
	TaskPriorityMedium TaskPriority = "MEDIUM"
	TaskPriorityHigh   TaskPriority = "HIGH"
)

type Task struct {
	ID          uint         `json:"id" gorm:"primaryKey"`
	Title       string       `json:"title" gorm:"not null"`
	Description string       `json:"description"`
	Status      TaskStatus   `json:"status" gorm:"type:varchar(20);default:'TODO'"`
	Priority    TaskPriority `json:"priority" gorm:"type:varchar(20);default:'MEDIUM'"`
	AssigneeID  *uint        `json:"assignee_id"` // Pointer to allow null
	Assignee    *User        `json:"assignee" gorm:"foreignKey:AssigneeID"`
	ProjectID   uint         `json:"project_id" gorm:"not null"`
	Project     Project      `json:"project" gorm:"foreignKey:ProjectID"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}
