package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

// Database functions
func initDB(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	// Create tables
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

	_, err = db.Exec(query)
	return db, err
}

func getAllTasks(completed *bool) ([]Task, error) {
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

	var tasks []Task
	for rows.Next() {
		var t Task
		var desc sql.NullString
		err := rows.Scan(&t.ID, &t.Title, &desc, &t.Completed, &t.Priority, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			return nil, err
		}
		if desc.Valid {
			t.Description = &desc.String
		}
		tasks = append(tasks, t)
	}

	return tasks, rows.Err()
}

func getTaskByID(id string) (*Task, error) {
	query := "SELECT id, title, description, completed, priority, created_at, updated_at FROM tasks WHERE id = ?"
	var t Task
	var desc sql.NullString
	err := db.QueryRow(query, id).Scan(&t.ID, &t.Title, &desc, &t.Completed, &t.Priority, &t.CreatedAt, &t.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	if desc.Valid {
		t.Description = &desc.String
	}
	return &t, nil
}

func createTask(input map[string]interface{}) (*Task, error) {
	now := time.Now().Format(time.RFC3339)
	priority := "MEDIUM"
	if p, ok := input["priority"].(string); ok && p != "" {
		priority = p
	}

	desc := ""
	if d, ok := input["description"].(string); ok {
		desc = d
	}

	task := Task{
		ID:        uuid.New().String(),
		Title:     input["title"].(string),
		Completed: false,
		Priority:  priority,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if desc != "" {
		task.Description = &desc
	}

	_, err := db.Exec(
		"INSERT INTO tasks (id, title, description, completed, priority, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)",
		task.ID, task.Title, desc, task.Completed, task.Priority, task.CreatedAt, task.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}
	return &task, nil
}

func updateTask(id string, input map[string]interface{}) (*Task, error) {
	task, err := getTaskByID(id)
	if err != nil {
		return nil, err
	}
	if task == nil {
		return nil, fmt.Errorf("task not found")
	}

	if title, ok := input["title"].(string); ok {
		task.Title = title
	}
	if desc, ok := input["description"].(string); ok {
		task.Description = &desc
	}
	if completed, ok := input["completed"].(bool); ok {
		task.Completed = completed
	}
	if priority, ok := input["priority"].(string); ok {
		task.Priority = priority
	}

	task.UpdatedAt = time.Now().Format(time.RFC3339)

	desc := ""
	if task.Description != nil {
		desc = *task.Description
	}

	_, err = db.Exec(
		"UPDATE tasks SET title = ?, description = ?, completed = ?, priority = ?, updated_at = ? WHERE id = ?",
		task.Title, desc, task.Completed, task.Priority, task.UpdatedAt, task.ID,
	)

	if err != nil {
		return nil, err
	}
	return task, nil
}

func deleteTask(id string) error {
	_, err := db.Exec("DELETE FROM tasks WHERE id = ?", id)
	return err
}

// TODO TASK 1: Implement this function
func getTasksByPriority(priority string) ([]Task, error) {
	// TASK: Query tasks WHERE priority = ?, ORDER BY created_at DESC
	query := "SELECT id, title, description, completed, priority, created_at, updated_at FROM tasks WHERE priority=? ORDER BY created_at DESC"
	rows, err := db.Query(query, priority)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		var desc sql.NullString
		err := rows.Scan(&t.ID, &t.Title, &desc, &t.Completed, &t.Priority, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			return nil, err
		}
		if desc.Valid {
			t.Description = &desc.String
		}
		tasks = append(tasks, t)
	}

	return tasks, rows.Err()
}

// TODO TASK 2: Implement this function
func searchTasks(search string) ([]Task, error) {
	searchTerm := "%" + search + "%"
	query := "SELECT id, title, description, completed, priority, created_at, updated_at FROM tasks WHERE title LIKE ? OR description LIKE ? ORDER by created_at DESC"
	rows, err := db.Query(query, searchTerm, searchTerm)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var tasks []Task
	for rows.Next() {
		var t Task
		var desc sql.NullString
		err := rows.Scan(&t.ID, &t.Title, &desc, &t.Completed, &t.Priority, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			return nil, err
		}
		if desc.Valid {
			t.Description = &desc.String
		}
		tasks = append(tasks, t)
	}

	return tasks, rows.Err()
}

// TODO TASK 3: Implement this function
func getTaskStats() (map[string]int, error) {
	// TASK: Return map with total, completed, pending, highPriority counts
	var total, completed, pending, highPriority int
	err := db.QueryRow("SELECT COUNT(*) FROM tasks").Scan(&total)
	if err != nil {
		return nil, err
	}

	err = db.QueryRow("SELECT COUNT(*) FROM tasks WHERE completed = true").Scan(&completed)

	if err != nil {
		return nil, err
	}

	err = db.QueryRow("SELECT COUNT(*) FROM tasks WHERE completed = false").Scan(&pending)
	if err != nil {
		return nil, err
	}

	err = db.QueryRow("SELECT COUNT(*) FROM tasks WHERE priority = 'HIGH'").Scan(&highPriority)
	if err != nil {
		return nil, err
	}

	stats := map[string]int{
		"total":        total,
		"completed":    completed,
		"pending":      pending,
		"highPriority": highPriority,
	}

	return stats, err
}

// GraphQL Resolvers
func resolveTasks(p graphql.ResolveParams) (interface{}, error) {
	var completed *bool
	if c, ok := p.Args["completed"].(bool); ok {
		completed = &c
	}
	return getAllTasks(completed)
}

func resolveTask(p graphql.ResolveParams) (interface{}, error) {
	id := p.Args["id"].(string)
	return getTaskByID(id)
}

// TODO TASK 1: Implement this resolver
func resolveTasksByPriority(p graphql.ResolveParams) (interface{}, error) {
	priority := p.Args["priority"].(string)
	return getTasksByPriority(priority)
}

// TODO TASK 2: Implement this resolver
func resolveSearchTasks(p graphql.ResolveParams) (interface{}, error) {
	// TASK: Get query from args, call searchTasks, return results
	search := p.Args["query"].(string)
	return searchTasks(search)
}

// TODO TASK 3: Implement this resolver
func resolveTaskStats(p graphql.ResolveParams) (interface{}, error) {
	// TASK: Call getTaskStats, return map matching TaskStats type
	return getTaskStats()
}

func resolveCreateTask(p graphql.ResolveParams) (interface{}, error) {
	input := p.Args["input"].(map[string]interface{})
	return createTask(input)
}

func resolveUpdateTask(p graphql.ResolveParams) (interface{}, error) {
	id := p.Args["id"].(string)
	input := p.Args["input"].(map[string]interface{})
	return updateTask(id, input)
}

func resolveDeleteTask(p graphql.ResolveParams) (interface{}, error) {
	id := p.Args["id"].(string)
	return true, deleteTask(id)
}

// TODO TASK 4: Implement this resolver
func resolveToggleTaskComplete(p graphql.ResolveParams) (interface{}, error) {
	// TASK: Get task by ID, toggle completed status, save, return updated task
	return nil, fmt.Errorf("TASK 4: Not implemented")
}
