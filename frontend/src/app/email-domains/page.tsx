'use client'

import { useState, useEffect } from 'react'
import { emailDomainsApi } from '@/lib/api'
import type { EmailDomain } from '@/types/api'

export default function EmailDomainsPage() {
  const [domains, setDomains] = useState<EmailDomain[]>([])
  const [loading, setLoading] = useState(true)
  const [showForm, setShowForm] = useState(false)
  const [formData, setFormData] = useState({
    domain: '',
    priority: 50,
    source: 'generator' as const,
  })

  useEffect(() => {
    loadDomains()
  }, [])

  async function loadDomains() {
    try {
      setLoading(true)
      const response = await emailDomainsApi.list()
      setDomains(response.data.domains || [])
    } catch (err) {
      console.error('Failed to load domains')
    } finally {
      setLoading(false)
    }
  }

  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault()
    try {
      await emailDomainsApi.create(formData)
      setShowForm(false)
      setFormData({ domain: '', priority: 50, source: 'generator' })
      loadDomains()
    } catch (err) {
      alert('Failed to create domain')
    }
  }

  async function handleToggle(id: string, isActive: boolean) {
    try {
      await emailDomainsApi.update(id, { is_active: !isActive })
      loadDomains()
    } catch (err) {
      alert('Failed to update domain')
    }
  }

  async function handleDelete(id: string) {
    if (!confirm('Are you sure you want to delete this domain?')) return
    try {
      await emailDomainsApi.delete(id)
      setDomains(domains.filter(d => d.id !== id))
    } catch (err) {
      alert('Failed to delete domain')
    }
  }

  async function handleHealthCheck(id: string) {
    try {
      const response = await emailDomainsApi.checkHealth(id)
      alert(`Health status: ${response.data.health_status}`)
      loadDomains()
    } catch (err) {
      alert('Failed to check health')
    }
  }

  if (loading) {
    return <div className="min-h-screen p-8"><div className="max-w-7xl mx-auto"><p>Loading...</p></div></div>
  }

  return (
    <div className="min-h-screen p-8">
      <div className="max-w-7xl mx-auto">
        <div className="flex justify-between items-center mb-6">
          <h1 className="text-3xl font-bold">Email Domains</h1>
          <button
            onClick={() => setShowForm(!showForm)}
            className="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600"
          >
            {showForm ? 'Cancel' : 'Add Domain'}
          </button>
        </div>

        {showForm && (
          <form onSubmit={handleSubmit} className="mb-6 p-6 bg-white rounded-lg shadow">
            <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">Domain</label>
                <input
                  type="text"
                  value={formData.domain}
                  onChange={(e) => setFormData({ ...formData, domain: e.target.value })}
                  className="w-full px-3 py-2 border border-gray-300 rounded-md"
                  placeholder="example.com"
                  required
                />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">Priority (1-100)</label>
                <input
                  type="number"
                  value={formData.priority}
                  onChange={(e) => setFormData({ ...formData, priority: parseInt(e.target.value) })}
                  className="w-full px-3 py-2 border border-gray-300 rounded-md"
                  min={1}
                  max={100}
                  required
                />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">Source</label>
                <select
                  value={formData.source}
                  onChange={(e) => setFormData({ ...formData, source: e.target.value as 'generator' | 'custom' })}
                  className="w-full px-3 py-2 border border-gray-300 rounded-md"
                >
                  <option value="generator">Generator.email</option>
                  <option value="custom">Custom</option>
                </select>
              </div>
            </div>
            <div className="mt-4">
              <button type="submit" className="px-4 py-2 bg-green-500 text-white rounded hover:bg-green-600">
                Create Domain
              </button>
            </div>
          </form>
        )}

        <div className="bg-white rounded-lg shadow overflow-hidden">
          <table className="min-w-full divide-y divide-gray-200">
            <thead className="bg-gray-50">
              <tr>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Domain</th>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Priority</th>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Active</th>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Health</th>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Source</th>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Actions</th>
              </tr>
            </thead>
            <tbody className="bg-white divide-y divide-gray-200">
              {domains.length === 0 ? (
                <tr>
                  <td colSpan={6} className="px-6 py-4 text-center text-gray-500">No domains found</td>
                </tr>
              ) : (
                domains.map((domain) => (
                  <tr key={domain.id}>
                    <td className="px-6 py-4 whitespace-nowrap text-sm font-medium">{domain.domain}</td>
                    <td className="px-6 py-4 whitespace-nowrap text-sm">{domain.priority}</td>
                    <td className="px-6 py-4 whitespace-nowrap text-sm">
                      <button
                        onClick={() => handleToggle(domain.id, domain.is_active)}
                        className={`px-2 py-1 rounded ${domain.is_active ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800'}`}
                      >
                        {domain.is_active ? 'Active' : 'Inactive'}
                      </button>
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap text-sm">
                      <span className={`px-2 py-1 rounded ${
                        domain.health_status === 'healthy' ? 'bg-green-100 text-green-800' :
                        domain.health_status === 'unhealthy' ? 'bg-red-100 text-red-800' :
                        'bg-gray-100 text-gray-800'
                      }`}>
                        {domain.health_status}
                      </span>
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap text-sm">{domain.source}</td>
                    <td className="px-6 py-4 whitespace-nowrap text-sm space-x-2">
                      <button onClick={() => handleHealthCheck(domain.id)} className="text-blue-600 hover:text-blue-900">Check</button>
                      <button onClick={() => handleDelete(domain.id)} className="text-red-600 hover:text-red-900">Delete</button>
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
