import { useState } from 'react'
import TaskList from './components/TaskList'
import TaskForm from './components/TaskForm'
import TaskStats from './components/TaskStats'
import PriorityFilter from './components/PriorityFilter'
import './App.css'

function App() {
  const [filter, setFilter] = useState<'all' | 'completed' | 'pending'>('all')

  return (
    <div className="min-h-screen bg-gray-100 py-8 px-4">
      <div className="max-w-4xl mx-auto">
        <header className="mb-8">
          <h1 className="text-3xl font-bold text-gray-800 mb-2">Task Manager</h1>
          <p className="text-gray-600">A GraphQL learning project</p>
        </header>

        <div className="grid gap-6">
          {/* Task Statistics - TODO TASK 7 */}
          <TaskStats />

          {/* Priority Filter Test - TASK 1/5 */}
          <PriorityFilter />

          {/* Create Task Form */}
          <TaskForm />

          {/* Filter Tabs */}
          <div className="flex gap-2">
            {(['all', 'pending', 'completed'] as const).map((f) => (
              <button
                key={f}
                onClick={() => setFilter(f)}
                className={`px-4 py-2 rounded-lg font-medium transition-colors ${
                  filter === f
                    ? 'bg-blue-600 text-white'
                    : 'bg-white text-gray-700 hover:bg-gray-50'
                }`}
              >
                {f.charAt(0).toUpperCase() + f.slice(1)}
              </button>
            ))}
          </div>

          {/* Task List */}
          <TaskList filter={filter} />
        </div>
      </div>
    </div>
  )
}

export default App
