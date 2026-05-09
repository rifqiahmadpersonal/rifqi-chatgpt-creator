'use client'

import { useState, useEffect } from 'react'
import { batchJobsApi } from '@/lib/api'
import type { BatchJob } from '@/types/api'

export default function BatchJobsPage() {
  const [jobs, setJobs] = useState<BatchJob[]>([])
  const [loading, setLoading] = useState(true)
  const [showForm, setShowForm] = useState(false)
  const [formData, setFormData] = useState({
    target_count: 5,
    max_workers: 3,
    default_password: '',
    proxy: '',
  })

  useEffect(() => {
    loadJobs()
  }, [])

  async function loadJobs() {
    try {
      setLoading(true)
      const response = await batchJobsApi.list()
      setJobs(response.data.jobs || [])
    } catch (err) {
      console.error('Failed to load jobs')
    } finally {
      setLoading(false)
    }
  }

  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault()
    try {
      const response = await batchJobsApi.create(formData)
      const jobId = response.data.job?.id
      setShowForm(false)
      loadJobs()
      if (jobId) {
        if (confirm('Job created! Start it now?')) {
          await batchJobsApi.start(jobId)
          loadJobs()
        }
      }
    } catch (err) {
      alert('Failed to create job')
    }
  }

  async function handleStart(id: string) {
    try {
      await batchJobsApi.start(id)
      loadJobs()
    } catch (err) {
      alert('Failed to start job')
    }
  }

  async function handleStop(id: string) {
    try {
      await batchJobsApi.stop(id)
      loadJobs()
    } catch (err) {
      alert('Failed to stop job')
    }
  }

  if (loading) {
    return <div className="min-h-screen p-8"><div className="max-w-7xl mx-auto"><p>Loading...</p></div></div>
  }

  return (
    <div className="min-h-screen p-8">
      <div className="max-w-7xl mx-auto">
        <div className="flex justify-between items-center mb-6">
          <h1 className="text-3xl font-bold">Batch Jobs</h1>
          <button
            onClick={() => setShowForm(!showForm)}
            className="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600"
          >
            {showForm ? 'Cancel' : 'New Job'}
          </button>
        </div>

        {showForm && (
          <form onSubmit={handleSubmit} className="mb-6 p-6 bg-white rounded-lg shadow">
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">Target Accounts</label>
                <input
                  type="number"
                  value={formData.target_count}
                  onChange={(e) => setFormData({ ...formData, target_count: parseInt(e.target.value) })}
                  className="w-full px-3 py-2 border border-gray-300 rounded-md"
                  min={1}
                  required
                />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">Max Workers</label>
                <input
                  type="number"
                  value={formData.max_workers}
                  onChange={(e) => setFormData({ ...formData, max_workers: parseInt(e.target.value) })}
                  className="w-full px-3 py-2 border border-gray-300 rounded-md"
                  min={1}
                  max={20}
                  required
                />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">Default Password (optional, min 12 chars)</label>
                <input
                  type="text"
                  value={formData.default_password}
                  onChange={(e) => setFormData({ ...formData, default_password: e.target.value })}
                  className="w-full px-3 py-2 border border-gray-300 rounded-md"
                  placeholder="Leave empty for random"
                />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">Proxy (optional)</label>
                <input
                  type="text"
                  value={formData.proxy}
                  onChange={(e) => setFormData({ ...formData, proxy: e.target.value })}
                  className="w-full px-3 py-2 border border-gray-300 rounded-md"
                  placeholder="http://proxy:port"
                />
              </div>
            </div>
            <div className="mt-4">
              <button type="submit" className="px-4 py-2 bg-green-500 text-white rounded hover:bg-green-600">
                Create Job
              </button>
            </div>
          </form>
        )}

        <div className="bg-white rounded-lg shadow overflow-hidden">
          <table className="min-w-full divide-y divide-gray-200">
            <thead className="bg-gray-50">
              <tr>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">ID</th>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Status</th>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Progress</th>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Workers</th>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Created</th>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Actions</th>
              </tr>
            </thead>
            <tbody className="bg-white divide-y divide-gray-200">
              {jobs.length === 0 ? (
                <tr>
                  <td colSpan={6} className="px-6 py-4 text-center text-gray-500">No jobs found</td>
                </tr>
              ) : (
                jobs.map((job) => (
                  <tr key={job.id}>
                    <td className="px-6 py-4 whitespace-nowrap text-sm font-mono">{job.id.slice(0, 8)}...</td>
                    <td className="px-6 py-4 whitespace-nowrap text-sm">
                      <span className={`px-2 py-1 rounded ${
                        job.status === 'completed' ? 'bg-green-100 text-green-800' :
                        job.status === 'running' ? 'bg-blue-100 text-blue-800' :
                        job.status === 'failed' ? 'bg-red-100 text-red-800' :
                        job.status === 'cancelled' ? 'bg-yellow-100 text-yellow-800' :
                        'bg-gray-100 text-gray-800'
                      }`}>
                        {job.status}
                      </span>
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap text-sm">
                      <div className="flex items-center">
                        <span className="mr-2">{job.success_count}/{job.target_count}</span>
                        <div className="w-24 bg-gray-200 rounded-full h-2">
                          <div
                            className="bg-green-500 h-2 rounded-full"
                            style={{ width: `${(job.success_count / job.target_count) * 100}%` }}
                          />
                        </div>
                      </div>
                      {job.failure_count > 0 && (
                        <span className="text-red-500 text-xs">({job.failure_count} failed)</span>
                      )}
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap text-sm">{job.max_workers}</td>
                    <td className="px-6 py-4 whitespace-nowrap text-sm">
                      {new Date(job.created_at).toLocaleString()}
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap text-sm space-x-2">
                      {job.status === 'pending' && (
                        <button onClick={() => handleStart(job.id)} className="text-green-600 hover:text-green-900">Start</button>
                      )}
                      {job.status === 'running' && (
                        <button onClick={() => handleStop(job.id)} className="text-red-600 hover:text-red-900">Stop</button>
                      )}
                    </td>
                  </tr>
                ))
              )}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  )
}
