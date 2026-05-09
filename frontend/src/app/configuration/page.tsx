'use client'

import { useState, useEffect } from 'react'
import { configurationsApi } from '@/lib/api'
import type { Configuration } from '@/types/api'

const CONFIG_KEYS = [
  { key: 'default_proxy', label: 'Default Proxy', type: 'text' },
  { key: 'default_password', label: 'Default Password', type: 'password' },
  { key: 'default_domain', label: 'Default Domain', type: 'text' },
  { key: 'worker_pool_size', label: 'Worker Pool Size', type: 'number' },
  { key: 'max_retries', label: 'Max Retries', type: 'number' },
  { key: 'registration_timeout', label: 'Registration Timeout', type: 'text' },
]

export default function ConfigurationPage() {
  const [configs, setConfigs] = useState<Configuration[]>([])
  const [loading, setLoading] = useState(true)
  const [editingKey, setEditingKey] = useState<string | null>(null)
  const [editValue, setEditValue] = useState('')

  useEffect(() => {
    loadConfigs()
  }, [])

  async function loadConfigs() {
    try {
      setLoading(true)
      const response = await configurationsApi.list()
      setConfigs(response.data.configurations || [])
    } catch (err) {
      console.error('Failed to load configurations')
    } finally {
      setLoading(false)
    }
  }

  function getValue(key: string): string {
    const config = configs.find(c => c.key === key)
    return config?.value || ''
  }

  function startEdit(key: string) {
    setEditingKey(key)
    setEditValue(getValue(key))
  }

  async function saveEdit(key: string) {
    try {
      await configurationsApi.update(key, editValue)
      setEditingKey(null)
      loadConfigs()
    } catch (err) {
      alert('Failed to update configuration')
    }
  }

  function cancelEdit() {
    setEditingKey(null)
    setEditValue('')
  }

  if (loading) {
    return <div className="min-h-screen p-8"><div className="max-w-7xl mx-auto"><p>Loading...</p></div></div>
  }

  return (
    <div className="min-h-screen p-8">
      <div className="max-w-4xl mx-auto">
        <h1 className="text-3xl font-bold mb-6">Configuration</h1>

        <div className="bg-white rounded-lg shadow">
          <div className="divide-y divide-gray-200">
            {CONFIG_KEYS.map(({ key, label, type }) => (
              <div key={key} className="p-6">
                <div className="flex justify-between items-start">
                  <div className="flex-1">
                    <label className="block text-sm font-medium text-gray-700 mb-1">{label}</label>
                    {editingKey === key ? (
                      <div className="flex items-center space-x-2">
                        <input
                          type={type}
                          value={editValue}
                          onChange={(e) => setEditValue(e.target.value)}
                          className="flex-1 px-3 py-2 border border-gray-300 rounded-md"
                        />
                        <button
                          onClick={() => saveEdit(key)}
                          className="px-3 py-2 bg-green-500 text-white rounded hover:bg-green-600"
                        >
                          Save
                        </button>
                        <button
                          onClick={cancelEdit}
                          className="px-3 py-2 bg-gray-300 text-gray-700 rounded hover:bg-gray-400"
                        >
                          Cancel
                        </button>
                      </div>
                    ) : (
                      <div className="flex items-center justify-between">
                        <span className="text-gray-600">
                          {type === 'password' && getValue(key) ? '••••••••' : getValue(key) || '(not set)'}
                        </span>
                        <button
                          onClick={() => startEdit(key)}
                          className="text-blue-600 hover:text-blue-900 text-sm"
                        >
                          Edit
                        </button>
                      </div>
                    )}
                  </div>
                </div>
              </div>
            ))}
          </div>
        </div>
      </div>
    </div>
  )
}
