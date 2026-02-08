import { useState } from 'react'
import { useQuery } from '@apollo/client'
import { SEARCH_TASKS, GET_TASKS } from '../graphql/queries'
import { Task } from '../types'
import TaskItem from './TaskItem'

export default function SearchTasks() {
  const [searchQuery, setSearchQuery] = useState('')
  const [isSearching, setIsSearching] = useState(false)

  // Only run search query when user submits
  const { data, loading, error, refetch } = useQuery(SEARCH_TASKS, {
    variables: { query: searchQuery },
    skip: !isSearching || !searchQuery.trim(),
  })

  // Get all tasks when not searching
  const { data: allData } = useQuery(GET_TASKS, {
    skip: isSearching,
  })

  const tasks: Task[] = isSearching ? data?.searchTasks || [] : allData?.tasks || []

  const handleSearch = (e: React.FormEvent) => {
    e.preventDefault()
    if (searchQuery.trim()) {
      setIsSearching(true)
      refetch({ query: searchQuery })
    }
  }

  const handleClear = () => {
    setSearchQuery('')
    setIsSearching(false)
  }

  return (
    <div className="bg-white rounded-lg shadow p-6 mb-6">
      <h3 className="text-lg font-semibold text-gray-800 mb-4">
        Test: Search Tasks (Task 2)
      </h3>

      {/* Search Form */}
      <form onSubmit={handleSearch} className="flex gap-2 mb-4">
        <div className="relative flex-1">
          <svg className="absolute left-3 top-1/2 -translate-y-1/2 w-5 h-5 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
          </svg>
          <input
            type="text"
            value={searchQuery}
            onChange={(e) => setSearchQuery(e.target.value)}
            placeholder="Search tasks by title or description..."
            className="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          />
        </div>
        <button
          type="submit"
          disabled={!searchQuery.trim() || loading}
          className="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 disabled:bg-gray-400 disabled:cursor-not-allowed flex items-center gap-2"
        >
          {loading ? (
            <>
              <div className="w-4 h-4 border-2 border-white border-t-transparent rounded-full animate-spin" />
              Searching...
            </>
          ) : (
            <>
              <svg className="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
              </svg>
              Search
            </>
          )}
        </button>
        {isSearching && (
          <button
            type="button"
            onClick={handleClear}
            className="px-4 py-2 bg-gray-200 text-gray-700 rounded-lg hover:bg-gray-300 flex items-center gap-2"
          >
            <svg className="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M6 18L18 6M6 6l12 12" />
            </svg>
            Clear
          </button>
        )}
      </form>

      {/* Status */}
      {isSearching && (
        <p className="text-sm text-gray-600 mb-4">
          Searching for: <span className="font-semibold">&quot;{searchQuery}&quot;</span>
        </p>
      )}

      {/* Error */}
      {error && (
        <div className="bg-red-50 border border-red-200 rounded-lg p-4 mb-4">
          <p className="text-red-600 font-medium">Backend Error</p>
          <p className="text-red-500 text-sm mt-1">{error.message}</p>
          <p className="text-red-400 text-xs mt-2">
            💡 This is expected! You need to implement Task 2 in the backend:
            <code className="bg-red-100 px-1 py-0.5 rounded ml-1">searchTasks</code> function in{' '}
            <code className="bg-red-100 px-1 py-0.5 rounded">backend/resolvers.go</code>
          </p>
        </div>
      )}

      {/* Results */}
      {!loading && !error && (
        <>
          {tasks.length === 0 ? (
            <div className="text-center py-8 bg-gray-50 rounded-lg">
              <p className="text-gray-500">
                {isSearching 
                  ? `No tasks found matching "${searchQuery}"` 
                  : 'Enter a search term to find tasks'}
              </p>
            </div>
          ) : (
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
        </>
      )}

      {/* Debug Info */}
      {isSearching && (
        <div className="mt-4 p-3 bg-gray-50 rounded-lg text-xs text-gray-600 font-mono">
          <p><strong>Debug Info:</strong></p>
          <p>Query: SEARCH_TASKS</p>
          <p>Variables: {'{ query: "'}{searchQuery}{'" }'}</p>
          <p>Results: {tasks.length}</p>
          <p>Status: {loading ? 'LOADING' : error ? 'ERROR' : 'SUCCESS'}</p>
        </div>
      )}
    </div>
  )
}
