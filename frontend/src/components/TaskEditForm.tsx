import { useState } from 'react'
import { useMutation } from '@apollo/client'
import { UPDATE_TASK, GET_TASKS } from '../graphql/queries'
import { Task, Priority } from '../types'

interface TaskEditFormProps {
  task: Task
  onCancel: () => void
}

export default function TaskEditForm({ task, onCancel }: TaskEditFormProps) {
  const [title, setTitle] = useState(task.title)
  const [description, setDescription] = useState(task.description || '')
  const [priority, setPriority] = useState<Priority>(task.priority)

  const [updateTask, { loading }] = useMutation(UPDATE_TASK, {
    refetchQueries: [{ query: GET_TASKS }],
    onCompleted: onCancel,
  })

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    if (!title.trim()) return

    await updateTask({
      variables: {
        id: task.id,
        input: {
          title,
          description: description || undefined,
          priority,
        },
      },
    })
  }

  return (
    <div className="p-4 bg-blue-50">
      <form onSubmit={handleSubmit} className="space-y-3">
        <input
          type="text"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
          required
        />

        <textarea
          value={description}
          onChange={(e) => setDescription(e.target.value)}
          rows={2}
          className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
        />

        <div className="flex gap-2">
          <select
            value={priority}
            onChange={(e) => setPriority(e.target.value as Priority)}
            className="px-3 py-2 border border-gray-300 rounded-lg"
          >
            <option value={Priority.LOW}>Low</option>
            <option value={Priority.MEDIUM}>Medium</option>
            <option value={Priority.HIGH}>High</option>
          </select>

          <div className="flex gap-2 ml-auto">
            <button
              type="button"
              onClick={onCancel}
              className="px-4 py-2 text-gray-600 hover:bg-gray-100 rounded-lg"
            >
              Cancel
            </button>
            <button
              type="submit"
              disabled={loading}
              className="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 disabled:bg-gray-400"
            >
              {loading ? 'Saving...' : 'Save'}
            </button>
          </div>
        </div>
      </form>
    </div>
  )
}
