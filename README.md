# Go + GraphQL Full-Stack Starter

A learning project for Go, GraphQL, and React integration.

## Project Structure

```
go-graphql-starter/
├── backend/          # Go + gqlgen GraphQL API
└── frontend/         # React + TypeScript + Apollo Client
```

## Tech Stack

- **Backend:** Go, gqlgen (GraphQL), SQLite, go-chi (router)
- **Frontend:** React, TypeScript, Vite, Apollo Client, Tailwind CSS

## Quick Start

### Backend

```bash
cd backend
go mod download
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

See [GitHub Issues](../../issues) for tasks to complete.

## Schema Overview

```graphql
type Task {
  id: ID!
  title: String!
  description: String
  completed: Boolean!
  priority: Priority!
  createdAt: String!
}

enum Priority {
  LOW
  MEDIUM
  HIGH
}
```
