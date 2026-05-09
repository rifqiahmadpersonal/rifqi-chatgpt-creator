import axios from 'axios'
import type { Account, EmailDomain, BlacklistedDomain, Configuration, BatchJob, RegistrationAttempt, DashboardStats } from '@/types/api'

const API_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080/api'

const api = axios.create({
  baseURL: API_URL,
  headers: {
    'Content-Type': 'application/json',
  },
})

export const accountsApi = {
  list: () => api.get<{ accounts: Account[] }>('/accounts'),
  get: (id: string) => api.get<{ account: Account }>(`/accounts/${id}`),
  create: (data: Partial<Account>) => api.post<{ account: Account }>('/accounts', data),
  delete: (id: string) => api.delete(`/accounts/${id}`),
  export: () => api.get<{ accounts: Account[] }>('/accounts/export'),
}

export const emailDomainsApi = {
  list: () => api.get<{ domains: EmailDomain[] }>('/email-domains'),
  get: (id: string) => api.get<{ domain: EmailDomain }>(`/email-domains/${id}`),
  create: (data: Partial<EmailDomain>) => api.post<{ domain: EmailDomain }>('/email-domains', data),
  update: (id: string, data: Partial<EmailDomain>) => api.put<{ domain: EmailDomain }>(`/email-domains/${id}`, data),
  delete: (id: string) => api.delete(`/email-domains/${id}`),
  checkHealth: (id: string) => api.post<{ health_status: string }>(`/email-domains/${id}/check`),
}

export const batchJobsApi = {
  list: () => api.get<{ jobs: BatchJob[] }>('/batch-jobs'),
  get: (id: string) => api.get<{ job: BatchJob }>(`/batch-jobs/${id}`),
  create: (data: Partial<BatchJob>) => api.post<{ job: BatchJob }>('/batch-jobs', data),
  start: (id: string) => api.post<{ status: string }>(`/batch-jobs/${id}/start`),
  stop: (id: string) => api.post<{ status: string }>(`/batch-jobs/${id}/stop`),
  getAttempts: (id: string) => api.get<{ attempts: RegistrationAttempt[] }>(`/batch-jobs/${id}/attempts`),
}

export const configurationsApi = {
  list: () => api.get<{ configurations: Configuration[] }>('/configurations'),
  get: (key: string) => api.get<{ key: string; value: string }>(`/configurations/${key}`),
  update: (key: string, value: string) => api.put<{ key: string; value: string }>(`/configurations/${key}`, { value }),
}

export const blacklistedDomainsApi = {
  list: () => api.get<{ domains: BlacklistedDomain[] }>('/blacklisted-domains'),
  create: (data: Partial<BlacklistedDomain>) => api.post<{ domain: BlacklistedDomain }>('/blacklisted-domains', data),
  delete: (id: string) => api.delete(`/blacklisted-domains/${id}`),
}

export const statsApi = {
  dashboard: () => api.get<DashboardStats>('/stats/dashboard'),
}

export default api
