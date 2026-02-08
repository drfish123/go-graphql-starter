package graph

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"

	"backend/graph/model"
	"backend/internal/database"
)

// Resolver is the root resolver
type Resolver struct {
	DB *sql.DB
}

// Query returns the QueryResolver implementation
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

// Mutation returns the MutationResolver implementation
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

type queryResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }

// ============================================================
// QUERY RESOLVERS - READING DATA
// ============================================================

// Tasks returns all tasks (implemented ✓)
func (r *queryResolver) Tasks(ctx context.Context, completed *bool) ([]*model.Task, error) {
	dbTasks, err := database.GetAllTasks(r.DB, completed)
	if err != nil {
		return nil, err
	}

	return dbTasksToModels(dbTasks), nil
}

// Task returns a single task by ID (implemented ✓)
func (r *queryResolver) Task(ctx context.Context, id string) (*model.Task, error) {
	dbTask, err := database.GetTaskByID(r.DB, id)
	if err != nil {
		return nil, err
	}
	
	if dbTask == nil {
		return nil, nil
	}
	
	return dbTaskToModel(dbTask), nil
}

// TasksByPriority returns tasks filtered by priority (TODO: implement)
func (r *queryResolver) TasksByPriority(ctx context.Context, priority model.Priority) ([]*model.Task, error) {
	// TODO TASK 1: Implement this resolver
	// 1. Call database.GetTasksByPriority with the priority string
	// 2. Convert the returned []database.Task to []*model.Task
	// 3. Return the result
	
	panic("not implemented: TasksByPriority")
}

// SearchTasks searches tasks by title/description (TODO: implement)
func (r *queryResolver) SearchTasks(ctx context.Context, query string) ([]*model.Task, error) {
	// TODO TASK 2: Implement this resolver
	// 1. Call database.SearchTasks with the search query
	// 2. Convert the results to []*model.Task
	// 3. Return the result
	
	panic("not implemented: SearchTasks")
}

// TaskStats returns task statistics (TODO: implement)
func (r *queryResolver) TaskStats(ctx context.Context) (*model.TaskStats, error) {
	// TODO TASK 3: Implement this resolver
	// 1. Call database.GetTaskStats() to get the counts
	// 2. Create and return a *model.TaskStats with the values
	// Stats should include: total, completed, pending, highPriority
	
	panic("not implemented: TaskStats")
}

// ============================================================
// MUTATION RESOLVERS - MODIFYING DATA
// ============================================================

// CreateTask creates a new task (implemented ✓)
func (r *mutationResolver) CreateTask(ctx context.Context, input model.CreateTaskInput) (*model.Task, error) {
	now := time.Now().Format(time.RFC3339)
	priority := model.PriorityMedium
	
	if input.Priority != nil {
		priority = *input.Priority
	}
	
	desc := sql.NullString{}
	if input.Description != nil {
		desc.String = *input.Description
		desc.Valid = true
	}
	
	dbTask := &database.Task{
		ID:          uuid.New().String(),
		Title:       input.Title,
		Description: desc,
		Completed:   false,
		Priority:    string(priority),
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	
	if err := database.CreateTask(r.DB, dbTask); err != nil {
		return nil, err
	}
	
	return dbTaskToModel(dbTask), nil
}

// UpdateTask updates an existing task (implemented ✓)
func (r *mutationResolver) UpdateTask(ctx context.Context, id string, input model.UpdateTaskInput) (*model.Task, error) {
	// Get existing task
	dbTask, err := database.GetTaskByID(r.DB, id)
	if err != nil {
		return nil, err
	}
	
	if dbTask == nil {
		return nil, fmt.Errorf("task not found")
	}
	
	// Apply updates
	if input.Title != nil {
		dbTask.Title = *input.Title
	}
	if input.Description != nil {
		dbTask.Description = sql.NullString{
			String: *input.Description,
			Valid:  true,
		}
	}
	if input.Completed != nil {
		dbTask.Completed = *input.Completed
	}
	if input.Priority != nil {
		dbTask.Priority = string(*input.Priority)
	}
	
	dbTask.UpdatedAt = time.Now().Format(time.RFC3339)
	
	if err := database.UpdateTask(r.DB, dbTask); err != nil {
		return nil, err
	}
	
	return dbTaskToModel(dbTask), nil
}

// DeleteTask deletes a task by ID (implemented ✓)
func (r *mutationResolver) DeleteTask(ctx context.Context, id string) (bool, error) {
	err := database.DeleteTask(r.DB, id)
	if err != nil {
		return false, err
	}
	return true, nil
}

// ToggleTaskComplete toggles the completion status (TODO: implement)
func (r *mutationResolver) ToggleTaskComplete(ctx context.Context, id string) (*model.Task, error) {
	// TODO TASK 4: Implement this resolver
	// 1. Get the task by ID using database.GetTaskByID
	// 2. Toggle the Completed field (true -> false, false -> true)
	// 3. Update UpdatedAt to current time
	// 4. Save using database.UpdateTask
	// 5. Return the updated task
	
	panic("not implemented: ToggleTaskComplete")
}

// ============================================================
// HELPER FUNCTIONS
// ============================================================

func dbTaskToModel(dbTask *database.Task) *model.Task {
	task := &model.Task{
		ID:        dbTask.ID,
		Title:     dbTask.Title,
		Completed: dbTask.Completed,
		Priority:  model.Priority(dbTask.Priority),
		CreatedAt: dbTask.CreatedAt,
		UpdatedAt: dbTask.UpdatedAt,
	}
	
	if dbTask.Description.Valid {
		task.Description = &dbTask.Description.String
	}
	
	return task
}

func dbTasksToModels(dbTasks []database.Task) []*model.Task {
	tasks := make([]*model.Task, len(dbTasks))
	for i, dbTask := range dbTasks {
		tasks[i] = dbTaskToModel(&dbTask)
	}
	return tasks
}
