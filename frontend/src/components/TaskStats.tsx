import { useQuery } from '@apollo/client'
import { GET_TASK_STATS } from '../graphql/queries'

export default function TaskStats() {
  const { data, loading, error } = useQuery(GET_TASK_STATS)

  const stats = data?.taskStats

  if (loading) {
    return (
      <div className="bg-white rounded-lg shadow p-6">
        <h3 className="text-lg font-semibold text-gray-800 mb-4">Task Statistics</h3>
        <div className="flex items-center justify-center py-4">
          <div className="w-6 h-6 border-2 border-blue-600 border-t-transparent rounded-full animate-spin"></div>
        </div>
      </div>
    )
  }

  if (error) {
    return (
      <div className="bg-white rounded-lg shadow p-6">
        <h3 className="text-lg font-semibold text-gray-800 mb-4">Task Statistics</h3>
        <div className="bg-red-50 border border-red-200 rounded-lg p-4">
          <p className="text-red-600 font-medium">Backend Error</p>
          <p className="text-red-500 text-sm mt-1">{error.message}</p>
          <p className="text-red-400 text-xs mt-2">
            💡 You need to implement Task 3 in the backend:
            <code className="bg-red-100 px-1 py-0.5 rounded">getTaskStats</code> in{' '}
            <code className="bg-red-100 px-1 py-0.5 rounded">backend/resolvers.go</code>
          </p>
        </div>
      </div>
    )
  }

  return (
    <div className="bg-white rounded-lg shadow p-6">
      <h3 className="text-lg font-semibold text-gray-800 mb-4">Task Statistics</h3>
      <div className="grid grid-cols-4 gap-4">
        <div className="text-center p-4 bg-gray-50 rounded-lg">
          <p className="text-2xl font-bold text-gray-800">{stats?.total ?? 0}</p>
          <p className="text-sm text-gray-500">Total</p>
        </div>
        <div className="text-center p-4 bg-green-50 rounded-lg">
          <p className="text-2xl font-bold text-green-600">{stats?.completed ?? 0}</p>
          <p className="text-sm text-gray-600">Completed</p>
        </div>
        <div className="text-center p-4 bg-yellow-50 rounded-lg">
          <p className="text-2xl font-bold text-yellow-600">{stats?.pending ?? 0}</p>
          <p className="text-sm text-gray-600">Pending</p>
        </div>
        <div className="text-center p-4 bg-red-50 rounded-lg">
          <p className="text-2xl font-bold text-red-600">{stats?.highPriority ?? 0}</p>
          <p className="text-sm text-gray-600">High Priority</p>
        </div>
      </div>
    </div>
  )
}
