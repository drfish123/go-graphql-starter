import { gql } from '@apollo/client'

// ============================================================
// QUERIES - Reading data
// ============================================================

// Get all tasks (IMPLEMENTED ✓)
export const GET_TASKS = gql`
  query GetTasks($completed: Boolean) {
    tasks(completed: $completed) {
      id
      title
      description
      completed
      priority
      createdAt
      updatedAt
    }
  }
`

// Get a single task by ID (IMPLEMENTED ✓)
export const GET_TASK = gql`
  query GetTask($id: ID!) {
    task(id: $id) {
      id
      title
      description
      completed
      priority
      createdAt
      updatedAt
    }
  }
`

// TODO TASK 5: Write the query for TasksByPriority
// Should accept a priority parameter and return tasks with that priority
// export const GET_TASKS_BY_PRIORITY = gql`
//   query GetTasksByPriority($priority: Priority!) {
//     tasksByPriority(priority: $priority) {
//       id
//       title
//       description
//       completed
//       priority
//       createdAt
//       updatedAt
//     }
//   }
// `

// TODO TASK 6: Write the query for SearchTasks
// Should accept a search query string and return matching tasks
// export const SEARCH_TASKS = gql`
//   query SearchTasks($query: String!) {
//     searchTasks(query: $query) {
//       id
//       title
//       description
//       completed
//       priority
//       createdAt
//       updatedAt
//     }
//   }
// `

// TODO TASK 7: Write the query for TaskStats
// Should return the task statistics (total, completed, pending, highPriority)
// export const GET_TASK_STATS = gql`
//   query GetTaskStats {
//     taskStats {
//       total
//       completed
//       pending
//       highPriority
//     }
//   }
// `

// ============================================================
// MUTATIONS - Modifying data
// ============================================================

// Create a new task (IMPLEMENTED ✓)
export const CREATE_TASK = gql`
  mutation CreateTask($input: CreateTaskInput!) {
    createTask(input: $input) {
      id
      title
      description
      completed
      priority
      createdAt
      updatedAt
    }
  }
`

// Update an existing task (IMPLEMENTED ✓)
export const UPDATE_TASK = gql`
  mutation UpdateTask($id: ID!, $input: UpdateTaskInput!) {
    updateTask(id: $id, input: $input) {
      id
      title
      description
      completed
      priority
      updatedAt
    }
  }
`

// Delete a task (IMPLEMENTED ✓)
export const DELETE_TASK = gql`
  mutation DeleteTask($id: ID!) {
    deleteTask(id: $id)
  }
`

// TODO TASK 8: Write the mutation for ToggleTaskComplete
// Should toggle a task's completion status and return the updated task
// export const TOGGLE_TASK_COMPLETE = gql`
//   mutation ToggleTaskComplete($id: ID!) {
//     toggleTaskComplete(id: $id) {
//       id
//       title
//       description
//       completed
//       priority
//       createdAt
//       updatedAt
//     }
//   }
// `
