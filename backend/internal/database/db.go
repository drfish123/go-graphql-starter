package database

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

// Task represents a task in the database
type Task struct {
	ID          string
	Title       string
	Description sql.NullString
	Completed   bool
	Priority    string
	CreatedAt   string
	UpdatedAt   string
}

// InitDB initializes the SQLite database
func InitDB(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	// Create tables
	if err := createTables(db); err != nil {
		return nil, err
	}

	return db, nil
}

func createTables(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS tasks (
		id TEXT PRIMARY KEY,
		title TEXT NOT NULL,
		description TEXT,
		completed BOOLEAN DEFAULT FALSE,
		priority TEXT DEFAULT 'MEDIUM',
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`
	
	_, err := db.Exec(query)
	return err
}

// GetAllTasks retrieves all tasks from the database
func GetAllTasks(db *sql.DB, completed *bool) ([]Task, error) {
	query := "SELECT id, title, description, completed, priority, created_at, updated_at FROM tasks"
	var args []interface{}
	
	if completed != nil {
		query += " WHERE completed = ?"
		args = append(args, *completed)
	}
	
	query += " ORDER BY created_at DESC"
	
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanTasks(rows)
}

// GetTaskByID retrieves a single task by ID
func GetTaskByID(db *sql.DB, id string) (*Task, error) {
	query := "SELECT id, title, description, completed, priority, created_at, updated_at FROM tasks WHERE id = ?"
	
	var task Task
	err := db.QueryRow(query, id).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Completed,
		&task.Priority,
		&task.CreatedAt,
		&task.UpdatedAt,
	)
	
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	
	return &task, nil
}

// CreateTask creates a new task
func CreateTask(db *sql.DB, task *Task) error {
	query := `
		INSERT INTO tasks (id, title, description, completed, priority, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`
	
	_, err := db.Exec(query, task.ID, task.Title, task.Description, task.Completed, task.Priority, task.CreatedAt, task.UpdatedAt)
	return err
}

// UpdateTask updates an existing task
func UpdateTask(db *sql.DB, task *Task) error {
	query := `
		UPDATE tasks 
		SET title = ?, description = ?, completed = ?, priority = ?, updated_at = ?
		WHERE id = ?
	`
	
	_, err := db.Exec(query, task.Title, task.Description, task.Completed, task.Priority, task.UpdatedAt, task.ID)
	return err
}

// DeleteTask deletes a task by ID
func DeleteTask(db *sql.DB, id string) error {
	query := "DELETE FROM tasks WHERE id = ?"
	_, err := db.Exec(query, id)
	return err
}

// TODO: Implement GetTasksByPriority
func GetTasksByPriority(db *sql.DB, priority string) ([]Task, error) {
	// TASK: Implement this function
	// Should query tasks where priority = ? and return them
	return nil, fmt.Errorf("not implemented")
}

// TODO: Implement SearchTasks
func SearchTasks(db *sql.DB, searchQuery string) ([]Task, error) {
	// TASK: Implement this function
	// Should search title and description using LIKE %?%
	return nil, fmt.Errorf("not implemented")
}

// TODO: Implement GetTaskStats
func GetTaskStats(db *sql.DB) (map[string]int, error) {
	// TASK: Implement this function
	// Should return counts: total, completed, pending, highPriority
	return nil, fmt.Errorf("not implemented")
}

func scanTasks(rows *sql.Rows) ([]Task, error) {
	var tasks []Task
	
	for rows.Next() {
		var task Task
		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Completed,
			&task.Priority,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	
	return tasks, rows.Err()
}
