package model

// Task represents a task in the system
type Task struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description *string  `json:"description,omitempty"`
	Completed   bool     `json:"completed"`
	Priority    Priority `json:"priority"`
	CreatedAt   string   `json:"createdAt"`
	UpdatedAt   string   `json:"updatedAt"`
}

// Priority represents task priority levels
type Priority string

const (
	PriorityLow    Priority = "LOW"
	PriorityMedium Priority = "MEDIUM"
	PriorityHigh   Priority = "HIGH"
)

// CreateTaskInput input for creating tasks
type CreateTaskInput struct {
	Title       string   `json:"title"`
	Description *string  `json:"description,omitempty"`
	Priority    *Priority `json:"priority,omitempty"`
}

// UpdateTaskInput input for updating tasks
type UpdateTaskInput struct {
	Title       *string  `json:"title,omitempty"`
	Description *string  `json:"description,omitempty"`
	Completed   *bool    `json:"completed,omitempty"`
	Priority    *Priority `json:"priority,omitempty"`
}

// TaskStats represents task statistics
type TaskStats struct {
	Total        int `json:"total"`
	Completed    int `json:"completed"`
	Pending      int `json:"pending"`
	HighPriority int `json:"highPriority"`
}
