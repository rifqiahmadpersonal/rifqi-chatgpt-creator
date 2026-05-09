'use client'

import { useState, useEffect } from 'react'
import { statsApi } from '@/lib/api'
import type { DashboardStats } from '@/types/api'

export default function StatsPage() {
  const [stats, setStats] = useState<DashboardStats | null>(null)
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    loadStats()
    const interval = setInterval(loadStats, 30000)
    return () => clearInterval(interval)
  }, [])

  async function loadStats() {
    try {
      const response = await statsApi.dashboard()
      setStats(response.data)
    } catch (err) {
      console.error('Failed to load stats')
    } finally {
      setLoading(false)
    }
  }

  if (loading) {
    return <div className="min-h-screen p-8"><div className="max-w-7xl mx-auto"><p>Loading...</p></div></div>
  }

  return (
    <div className="min-h-screen p-8">
      <div className="max-w-7xl mx-auto">
        <h1 className="text-3xl font-bold mb-6">Dashboard Statistics</h1>

        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-5 gap-6">
          <StatCard title="Total Accounts" value={stats?.total_accounts || 0} color="blue" />
          <StatCard title="Active Accounts" value={stats?.active_accounts || 0} color="green" />
          <StatCard title="Total Batch Jobs" value={stats?.total_batch_jobs || 0} color="purple" />
          <StatCard title="Running Jobs" value={stats?.running_batch_jobs || 0} color="yellow" />
          <StatCard title="Active Domains" value={stats?.active_email_domains || 0} color="indigo" />
        </div>
      </div>
    </div>
  )
}

function StatCard({ title, value, color }: { title: string; value: number; color: string }) {
  const colorClasses: Record<string, string> = {
    blue: 'bg-blue-500',
    green: 'bg-green-500',
    purple: 'bg-purple-500',
    yellow: 'bg-yellow-500',
    indigo: 'bg-indigo-500',
  }

  return (
    <div className="bg-white rounded-lg shadow p-6">
      <div className={`w-12 h-12 ${colorClasses[color]} rounded-lg flex items-center justify-center mb-4`}>
        <span className="text-white text-2xl font-bold">{value}</span>
      </div>
      <h3 className="text-lg font-semibold text-gray-900">{value}</h3>
      <p className="text-sm text-gray-500">{title}</p>
    </div>
  )
}
