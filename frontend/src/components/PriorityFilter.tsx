import { useState } from 'react'
import { useQuery } from '@apollo/client'
import { GET_TASKS_BY_PRIORITY, GET_TASKS } from '../graphql/queries'
import { Priority, Task } from '../types'
import TaskItem from './TaskItem'

export default function PriorityFilter() {
  const [selectedPriority, setSelectedPriority] = useState<Priority | null>(null)

  // Use priority filter if selected, otherwise show all tasks
  const { data, loading, error } = useQuery(
    selectedPriority ? GET_TASKS_BY_PRIORITY : GET_TASKS,
    {
      variables: selectedPriority ? { priority: selectedPriority } : {},
      skip: false,
    }
  )

  const tasks: Task[] = data?.tasksByPriority || data?.tasks || []

  const priorityColors: Record<Priority, string> = {
    [Priority.LOW]: 'bg-gray-100 text-gray-800 border-gray-300',
    [Priority.MEDIUM]: 'bg-yellow-100 text-yellow-800 border-yellow-300',
    [Priority.HIGH]: 'bg-red-100 text-red-800 border-red-300',
  }

  return (
    <div className="bg-white rounded-lg shadow p-6 mb-6">
      <h3 className="text-lg font-semibold text-gray-800 mb-4">
        Test: Filter by Priority
      </h3>

      {/* Priority Buttons */}
      <div className="flex gap-2 mb-4">
        <button
          onClick={() => setSelectedPriority(null)}
          className={`px-4 py-2 rounded-lg border-2 font-medium transition-colors ${
            selectedPriority === null
              ? 'bg-blue-600 text-white border-blue-600'
              : 'bg-white text-gray-700 border-gray-300 hover:bg-gray-50'
          }`}
        >
          All Tasks
        </button>

        {(Object.values(Priority) as Priority[]).map((priority) => (
          <button
            key={priority}
            onClick={() => setSelectedPriority(priority)}
            className={`px-4 py-2 rounded-lg border-2 font-medium transition-colors ${
              selectedPriority === priority
                ? 'ring-2 ring-blue-500 ring-offset-2 ' + priorityColors[priority]
                : 'bg-white text-gray-700 border-gray-300 hover:bg-gray-50'
            }`}
          >
            {priority}
          </button>
        ))}
      </div>

      {/* Status */}
      {selectedPriority && (
        <p className="text-sm text-gray-600 mb-4">
          Showing tasks with priority: <span className="font-semibold">{selectedPriority}</span>
        </p>
      )}

      {/* Results */}
      {loading && (
        <div className="text-center py-4">
          <div className="inline-block animate-spin rounded-full h-6 w-6 border-b-2 border-blue-600"></div>
          <p className="mt-2 text-gray-600">Loading tasks...</p>
        </div>
      )}

      {error && (
        <div className="bg-red-50 border border-red-200 rounded-lg p-4">
          <p className="text-red-600">Error: {error.message}</p>
          <p className="text-sm text-red-500 mt-1">
            Make sure your backend implementation is working!
          </p>
        </div>
      )}

      {!loading && !error && tasks.length === 0 && (
        <p className="text-gray-500 text-center py-4">
          No tasks found with {selectedPriority || 'any'} priority.
        </p>
      )}

      {!loading && !error && tasks.length > 0 && (
        <div className="border rounded-lg divide-y divide-gray-100">
          {tasks.map((task) => (
            <TaskItem
              key={task.id}
              task={task}
              onDelete={() => {}}
            />
          ))}
        </div>
      )}

      {/* Debug Info */}
      {selectedPriority && (
        <div className="mt-4 p-3 bg-gray-50 rounded-lg text-xs text-gray-600">
          <p><strong>Debug Info:</strong></p>
          <p>Query: GET_TASKS_BY_PRIORITY</p>
          <p>Variables: {"{ priority: \"" + selectedPriority + "\" }"}</p>
          <p>Tasks returned: {tasks.length}</p>
        </div>
      )}
    </div>
  )
}
