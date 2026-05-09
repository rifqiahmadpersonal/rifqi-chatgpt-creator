import Link from 'next/link'

export default function Home() {
  return (
    <main className="min-h-screen p-8">
      <div className="max-w-7xl mx-auto">
        <header className="mb-8">
          <h1 className="text-4xl font-bold mb-2">ChatGPT Creator</h1>
          <p className="text-gray-600">Account Registration Bot Dashboard</p>
        </header>

        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mb-8">
          <DashboardCard
            title="Accounts"
            description="View and manage registered accounts"
            href="/accounts"
          />
          <DashboardCard
            title="Email Domains"
            description="Configure email domains for registration"
            href="/email-domains"
          />
          <DashboardCard
            title="Batch Jobs"
            description="Start and monitor batch registration jobs"
            href="/batch-jobs"
          />
          <DashboardCard
            title="Blacklisted Domains"
            description="View and manage blacklisted email domains"
            href="/blacklisted-domains"
          />
          <DashboardCard
            title="Configuration"
            description="Manage application settings"
            href="/configuration"
          />
          <DashboardCard
            title="Statistics"
            description="View registration statistics and metrics"
            href="/stats"
          />
        </div>
      </div>
    </main>
  )
}

function DashboardCard({ title, description, href }: { title: string; description: string; href: string }) {
  return (
    <Link
      href={href}
      className="block p-6 bg-white rounded-lg border border-gray-200 shadow-sm hover:shadow-md transition-shadow dark:bg-gray-800 dark:border-gray-700"
    >
      <h2 className="text-xl font-semibold mb-2">{title}</h2>
      <p className="text-gray-600 dark:text-gray-400">{description}</p>
    </Link>
  )
}
