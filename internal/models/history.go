package models

import "time"

type TaskHistory struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	TaskID        uint      `json:"task_id" gorm:"not null"`
	UserID        uint      `json:"user_id" gorm:"not null"`     // Who made the change
	ChangeType    string    `json:"change_type" gorm:"not null"` // e.g., "UPDATE", "CREATE", "DELETE"
	PreviousValue string    `json:"previous_value"`              // JSON string or simple text
	NewValue      string    `json:"new_value"`                   // JSON string or simple text
	CreatedAt     time.Time `json:"created_at"`
}
