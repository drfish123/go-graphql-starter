// TODO TASK 7: Implement TaskStats component
// 
// This component should:
// 1. Use the GET_TASK_STATS query to fetch statistics
// 2. Display a summary of tasks (total, completed, pending, high priority)
// 3. Show as a nice dashboard/card layout
//
// Hint: Import GET_TASK_STATS from '../graphql/queries' and use useQuery hook

export default function TaskStats() {
  // TODO: Implement this component
  // For now, just show a placeholder

  return (
    <div className="bg-white rounded-lg shadow p-6">
      <h3 className="text-lg font-semibold text-gray-800 mb-4">Task Statistics</h3>
      <div className="grid grid-cols-4 gap-4">
        <div className="text-center p-4 bg-gray-50 rounded-lg">
          <p className="text-2xl font-bold text-gray-400">?</p>
          <p className="text-sm text-gray-500">Total</p>
        </div>
        <div className="text-center p-4 bg-green-50 rounded-lg">
          <p className="text-2xl font-bold text-green-600">?</p>
          <p className="text-sm text-gray-600">Completed</p>
        </div>
        <div className="text-center p-4 bg-yellow-50 rounded-lg">
          <p className="text-2xl font-bold text-yellow-600">?</p>
          <p className="text-sm text-gray-600">Pending</p>
        </div>
        <div className="text-center p-4 bg-red-50 rounded-lg">
          <p className="text-2xl font-bold text-red-600">?</p>
          <p className="text-sm text-gray-600">High Priority</p>
        </div>
      </div>
      
      <p className="mt-4 text-sm text-gray-500 text-center">
        📋 Complete Task 7 to implement statistics
      </p>
    </div>
  )
}
