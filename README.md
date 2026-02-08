# Go + GraphQL Full-Stack Starter

A learning project for Go, GraphQL, and React integration.

## Project Structure

```
go-graphql-starter/
├── backend/
│   ├── server.go     # GraphQL schema setup and server
│   ├── resolvers.go  # Database functions and resolvers
│   └── go.mod
└── frontend/         # React + TypeScript + Apollo Client
```

## Tech Stack

- **Backend:** Go, graphql-go, SQLite, go-chi (router)
- **Frontend:** React, TypeScript, Vite, Apollo Client, Tailwind CSS

## Quick Start

### Backend

```bash
cd backend
go mod tidy
go run server.go
```

GraphQL Playground: http://localhost:8080

### Frontend

```bash
cd frontend
npm install
npm run dev
```

App: http://localhost:5173

## Learning Tasks

The codebase contains `TODO TASK` comments. Each task has a corresponding GitHub issue:

| Task | Description | File |
|------|-------------|------|
| #1 | Implement `getTasksByPriority` | `backend/resolvers.go` |
| #2 | Implement `searchTasks` | `backend/resolvers.go` |
| #3 | Implement `getTaskStats` | `backend/resolvers.go` |
| #4 | Implement `resolveToggleTaskComplete` | `backend/resolvers.go` |
| #5 | Write `GET_TASKS_BY_PRIORITY` query | `frontend/src/graphql/queries.ts` |
| #6 | Write `SEARCH_TASKS` query | `frontend/src/graphql/queries.ts` |
| #7 | Implement `TaskStats` component | `frontend/src/components/TaskStats.tsx` |
| #8 | Implement `TOGGLE_TASK_COMPLETE` mutation | `frontend/src/graphql/queries.ts` + `TaskItem.tsx` |

## Schema

```graphql
type Task {
  id: ID!
  title: String!
  description: String
  completed: Boolean!
  priority: Priority!
  createdAt: String!
  updatedAt: String!
}

enum Priority {
  LOW
  MEDIUM
  HIGH
}

type TaskStats {
  total: Int!
  completed: Int!
  pending: Int!
  highPriority: Int!
}

type Query {
  tasks(completed: Boolean): [Task!]!
  task(id: ID!): Task
  tasksByPriority(priority: Priority!): [Task!]!
  searchTasks(query: String!): [Task!]!
  taskStats: TaskStats!
}

type Mutation {
  createTask(input: CreateTaskInput!): Task!
  updateTask(id: ID!, input: UpdateTaskInput!): Task!
  deleteTask(id: ID!): Boolean!
  toggleTaskComplete(id: ID!): Task!
}
```

## Working Features (No Tasks Needed)

- ✅ Create tasks
- ✅ Update tasks
- ✅ Delete tasks
- ✅ List all tasks with filter
- ✅ Get single task by ID

## Tips

1. Start with **Task 1** (backend) - it's the easiest
2. Use the GraphQL Playground to test your backend changes
3. Each resolver function has a `// TASK X` comment and hints
4. Look at working functions like `getAllTasks` for patterns to follow
