import { useState } from 'react'
import { Task, Priority } from '../types'
import TaskEditForm from './TaskEditForm'
import { useMutation } from '@apollo/client'
import { GET_TASKS, TOGGLE_TASK_COMPLETE } from '../graphql/queries'

interface TaskItemProps {
  task: Task
  onDelete: (id: string) => void
}

export default function TaskItem({ task, onDelete }: TaskItemProps) {
  const [isEditing, setIsEditing] = useState(false)

  const [toggleTaskComplete] = useMutation(TOGGLE_TASK_COMPLETE, {
      refetchQueries: [{ query: GET_TASKS }],
  })

  // TODO TASK 8: Implement toggle functionality
  // This should call the TOGGLE_TASK_COMPLETE mutation
  const handleToggle = async () => {
    
      toggleTaskComplete({
        variables:{
          id:task.id
        }
      })
  }

  const priorityColors: Record<Priority, string> = {
    [Priority.LOW]: 'bg-gray-100 text-gray-800',
    [Priority.MEDIUM]: 'bg-yellow-100 text-yellow-800',
    [Priority.HIGH]: 'bg-red-100 text-red-800',
  }

  if (isEditing) {
    return (
      <TaskEditForm
        task={task}
        onCancel={() => setIsEditing(false)}
      />
    )
  }

  return (
    <div className={`p-4 flex items-center gap-4 ${task.completed ? 'bg-gray-50' : ''}`}>
      <input
        type="checkbox"
        checked={task.completed}
        onChange={handleToggle}
        className="w-5 h-5 rounded border-gray-300 text-blue-600 focus:ring-blue-500"
      />

      <div className="flex-1 min-w-0">
        <p className={`font-medium ${task.completed ? 'line-through text-gray-400' : 'text-gray-900'}`}>
          {task.title}
        </p>
        {task.description && (
          <p className="text-sm text-gray-500 truncate">{task.description}</p>
        )}
        <div className="flex items-center gap-2 mt-1">
          <span className={`text-xs px-2 py-0.5 rounded-full font-medium ${priorityColors[task.priority]}`}>
            {task.priority}
          </span>
          <span className="text-xs text-gray-400">
            {new Date(task.createdAt).toLocaleDateString()}
          </span>
        </div>
      </div>

      <div className="flex items-center gap-2">
        <button
          onClick={() => setIsEditing(true)}
          className="px-3 py-1 text-sm text-blue-600 hover:bg-blue-50 rounded transition-colors"
        >
          Edit
        </button>
        <button
          onClick={() => onDelete(task.id)}
          className="px-3 py-1 text-sm text-red-600 hover:bg-red-50 rounded transition-colors"
        >
          Delete
        </button>
      </div>
    </div>
  )
}
