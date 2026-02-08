package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	_ "modernc.org/sqlite"
)

// Task represents a task
type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description *string `json:"description,omitempty"`
	Completed   bool   `json:"completed"`
	Priority    string `json:"priority"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

var db *sql.DB

func main() {
	var err error
	db, err = initDB("tasks.db")
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	// Setup GraphQL schema
	schema, err := setupGraphQL()
	if err != nil {
		log.Fatal("Failed to setup GraphQL:", err)
	}

	// Create router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(corsMiddleware)

	// GraphQL handler
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	r.Handle("/", h)
	r.Handle("/query", h)

	log.Println("Server starting on http://localhost:8080")
	log.Println("GraphQL Playground: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func setupGraphQL() (graphql.Schema, error) {
	// Priority enum
	priorityEnum := graphql.NewEnum(graphql.EnumConfig{
		Name: "Priority",
		Values: graphql.EnumValueConfigMap{
			"LOW":    &graphql.EnumValueConfig{Value: "LOW"},
			"MEDIUM": &graphql.EnumValueConfig{Value: "MEDIUM"},
			"HIGH":   &graphql.EnumValueConfig{Value: "HIGH"},
		},
	})

	// Task type
	taskType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Task",
		Fields: graphql.Fields{
			"id":          &graphql.Field{Type: graphql.NewNonNull(graphql.ID)},
			"title":       &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			"description": &graphql.Field{Type: graphql.String},
			"completed":   &graphql.Field{Type: graphql.NewNonNull(graphql.Boolean)},
			"priority":    &graphql.Field{Type: graphql.NewNonNull(priorityEnum)},
			"createdAt":   &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			"updatedAt":   &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		},
	})

	// TaskStats type
	taskStatsType := graphql.NewObject(graphql.ObjectConfig{
		Name: "TaskStats",
		Fields: graphql.Fields{
			"total":        &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
			"completed":    &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
			"pending":      &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
			"highPriority": &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
		},
	})

	// Input types
	createTaskInput := graphql.NewInputObject(graphql.InputObjectConfig{
		Name: "CreateTaskInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"title":       &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
			"description": &graphql.InputObjectFieldConfig{Type: graphql.String},
			"priority":    &graphql.InputObjectFieldConfig{Type: priorityEnum},
		},
	})

	updateTaskInput := graphql.NewInputObject(graphql.InputObjectConfig{
		Name: "UpdateTaskInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"title":       &graphql.InputObjectFieldConfig{Type: graphql.String},
			"description": &graphql.InputObjectFieldConfig{Type: graphql.String},
			"completed":   &graphql.InputObjectFieldConfig{Type: graphql.Boolean},
			"priority":    &graphql.InputObjectFieldConfig{Type: priorityEnum},
		},
	})

	// Root Query
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"tasks": &graphql.Field{
				Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(taskType))),
				Args: graphql.FieldConfigArgument{
					"completed": &graphql.ArgumentConfig{Type: graphql.Boolean},
				},
				Resolve: resolveTasks,
			},
			"task": &graphql.Field{
				Type: taskType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.ID)},
				},
				Resolve: resolveTask,
			},
			"tasksByPriority": &graphql.Field{
				Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(taskType))),
				Args: graphql.FieldConfigArgument{
					"priority": &graphql.ArgumentConfig{Type: graphql.NewNonNull(priorityEnum)},
				},
				Resolve: resolveTasksByPriority,
			},
			"searchTasks": &graphql.Field{
				Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(taskType))),
				Args: graphql.FieldConfigArgument{
					"query": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				},
				Resolve: resolveSearchTasks,
			},
			"taskStats": &graphql.Field{
				Type: graphql.NewNonNull(taskStatsType),
				Resolve: resolveTaskStats,
			},
		},
	})

	// Root Mutation
	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createTask": &graphql.Field{
				Type: graphql.NewNonNull(taskType),
				Args: graphql.FieldConfigArgument{
					"input": &graphql.ArgumentConfig{Type: graphql.NewNonNull(createTaskInput)},
				},
				Resolve: resolveCreateTask,
			},
			"updateTask": &graphql.Field{
				Type: graphql.NewNonNull(taskType),
				Args: graphql.FieldConfigArgument{
					"id":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.ID)},
					"input": &graphql.ArgumentConfig{Type: graphql.NewNonNull(updateTaskInput)},
				},
				Resolve: resolveUpdateTask,
			},
			"deleteTask": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Boolean),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.ID)},
				},
				Resolve: resolveDeleteTask,
			},
			"toggleTaskComplete": &graphql.Field{
				Type: graphql.NewNonNull(taskType),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.ID)},
				},
				Resolve: resolveToggleTaskComplete,
			},
		},
	})

	return graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})
}
