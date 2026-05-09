export interface Account {
  id: string
  email: string
  status: 'active' | 'inactive' | 'suspended'
  batch_job_id?: string
  created_at: string
  updated_at: string
}

export interface EmailDomain {
  id: string
  domain: string
  priority: number
  is_active: boolean
  source: 'generator' | 'custom'
  last_checked?: string
  health_status: 'healthy' | 'unhealthy' | 'unknown'
  created_at: string
  updated_at: string
}

export interface BlacklistedDomain {
  id: string
  domain: string
  reason: string
  created_at: string
}

export interface Configuration {
  id: string
  key: string
  value: string
  created_at: string
  updated_at: string
}

export interface BatchJob {
  id: string
  target_count: number
  success_count: number
  failure_count: number
  status: 'pending' | 'running' | 'completed' | 'cancelled' | 'failed'
  max_workers: number
  default_password?: string
  proxy?: string
  created_at: string
  completed_at?: string
}

export interface RegistrationAttempt {
  id: string
  email: string
  status: 'success' | 'failed' | 'in_progress'
  error_message?: string
  worker_id: number
  batch_job_id: string
  started_at: string
  completed_at?: string
  duration_ms?: number
}

export interface DashboardStats {
  total_accounts: number
  active_accounts: number
  total_batch_jobs: number
  running_batch_jobs: number
  active_email_domains: number
}
