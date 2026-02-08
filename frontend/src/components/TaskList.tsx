import { useQuery, useMutation } from '@apollo/client'
import { GET_TASKS, DELETE_TASK } from '../graphql/queries'
import { Task } from '../types'
import TaskItem from './TaskItem'

interface TaskListProps {
  filter: 'all' | 'completed' | 'pending'
}

export default function TaskList({ filter }: TaskListProps) {
  // Convert filter to completed parameter
  const completedParam = filter === 'all' ? undefined : filter === 'completed'

  const { data, loading, error } = useQuery(GET_TASKS, {
    variables: { completed: completedParam },
  })

  const [deleteTask] = useMutation(DELETE_TASK, {
    refetchQueries: [{ query: GET_TASKS, variables: { completed: completedParam } }],
  })

  const handleDelete = async (id: string) => {
    if (confirm('Are you sure you want to delete this task?')) {
      await deleteTask({ variables: { id } })
    }
  }

  if (loading) {
    return (
      <div className="bg-white rounded-lg shadow p-8 text-center">
        <div className="inline-block animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"></div>
        <p className="mt-2 text-gray-600">Loading tasks...</p>
      </div>
    )
  }

  if (error) {
    return (
      <div className="bg-red-50 border border-red-200 rounded-lg p-4">
        <p className="text-red-600">Error loading tasks: {error.message}</p>
      </div>
    )
  }

  const tasks: Task[] = data?.tasks || []

  if (tasks.length === 0) {
    return (
      <div className="bg-white rounded-lg shadow p-8 text-center">
        <p className="text-gray-500">No tasks found. Create one above!</p>
      </div>
    )
  }

  return (
    <div className="bg-white rounded-lg shadow divide-y divide-gray-100">
      {tasks.map((task) => (
        <TaskItem
          key={task.id}
          task={task}
          onDelete={handleDelete}
        />
      ))}
    </div>
  )
}
