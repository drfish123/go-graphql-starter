// Task Priority Levels
export enum Priority {
  LOW = 'LOW',
  MEDIUM = 'MEDIUM',
  HIGH = 'HIGH',
}

// Task type
export interface Task {
  id: string
  title: string
  description?: string
  completed: boolean
  priority: Priority
  createdAt: string
  updatedAt: string
}

// Task Statistics
export interface TaskStats {
  total: number
  completed: number
  pending: number
  highPriority: number
}

// Input for creating a task
export interface CreateTaskInput {
  title: string
  description?: string
  priority?: Priority
}

// Input for updating a task
export interface UpdateTaskInput {
  title?: string
  description?: string
  completed?: boolean
  priority?: Priority
}
