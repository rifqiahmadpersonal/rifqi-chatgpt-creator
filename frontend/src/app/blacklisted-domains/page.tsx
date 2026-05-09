'use client'

import { useState, useEffect } from 'react'
import { blacklistedDomainsApi } from '@/lib/api'
import type { BlacklistedDomain } from '@/types/api'

export default function BlacklistedDomainsPage() {
  const [domains, setDomains] = useState<BlacklistedDomain[]>([])
  const [loading, setLoading] = useState(true)
  const [showForm, setShowForm] = useState(false)
  const [formData, setFormData] = useState({ domain: '', reason: '' })

  useEffect(() => {
    loadDomains()
  }, [])

  async function loadDomains() {
    try {
      setLoading(true)
      const response = await blacklistedDomainsApi.list()
      setDomains(response.data.domains || [])
    } catch (err) {
      console.error('Failed to load blacklisted domains')
    } finally {
      setLoading(false)
    }
  }

  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault()
    try {
      await blacklistedDomainsApi.create(formData)
      setShowForm(false)
      setFormData({ domain: '', reason: '' })
      loadDomains()
    } catch (err) {
      alert('Failed to add domain')
    }
  }

  async function handleDelete(id: string) {
    if (!confirm('Are you sure you want to remove this domain from the blacklist?')) return
    try {
      await blacklistedDomainsApi.delete(id)
      setDomains(domains.filter(d => d.id !== id))
    } catch (err) {
      alert('Failed to delete domain')
    }
  }

  if (loading) {
    return <div className="min-h-screen p-8"><div className="max-w-7xl mx-auto"><p>Loading...</p></div></div>
  }

  return (
    <div className="min-h-screen p-8">
      <div className="max-w-7xl mx-auto">
        <div className="flex justify-between items-center mb-6">
          <h1 className="text-3xl font-bold">Blacklisted Domains</h1>
          <button
            onClick={() => setShowForm(!showForm)}
            className="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600"
          >
            {showForm ? 'Cancel' : 'Add Domain'}
          </button>
        </div>

        {showForm && (
          <form onSubmit={handleSubmit} className="mb-6 p-6 bg-white rounded-lg shadow">
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
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
                <label className="block text-sm font-medium text-gray-700 mb-1">Reason</label>
                <input
                  type="text"
                  value={formData.reason}
                  onChange={(e) => setFormData({ ...formData, reason: e.target.value })}
                  className="w-full px-3 py-2 border border-gray-300 rounded-md"
                  placeholder="Unsupported email"
                />
              </div>
            </div>
            <div className="mt-4">
              <button type="submit" className="px-4 py-2 bg-green-500 text-white rounded hover:bg-green-600">
                Add to Blacklist
              </button>
            </div>
          </form>
        )}

        <div className="bg-white rounded-lg shadow overflow-hidden">
          <table className="min-w-full divide-y divide-gray-200">
            <thead className="bg-gray-50">
              <tr>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Domain</th>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Reason</th>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Added</th>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Actions</th>
              </tr>
            </thead>
            <tbody className="bg-white divide-y divide-gray-200">
              {domains.length === 0 ? (
                <tr>
                  <td colSpan={4} className="px-6 py-4 text-center text-gray-500">No blacklisted domains</td>
                </tr>
              ) : (
                domains.map((domain) => (
                  <tr key={domain.id}>
                    <td className="px-6 py-4 whitespace-nowrap text-sm font-medium">{domain.domain}</td>
                    <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{domain.reason || '-'}</td>
                    <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                      {new Date(domain.created_at).toLocaleDateString()}
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap text-sm">
                      <button
                        onClick={() => handleDelete(domain.id)}
                        className="text-red-600 hover:text-red-900"
                      >
                        Remove
                      </button>
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
