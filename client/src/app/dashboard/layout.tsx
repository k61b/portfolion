import Navbar from '@components/navbar'

export const metadata = {
  title: 'Dashboard',
  description: 'Dashboard for bookmarks',
}

export default function DashboardLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <section className="flex flex-col h-screen">
      <Navbar />
      <main className="flex flex-col flex-grow">{children}</main>
    </section>
  )
}
