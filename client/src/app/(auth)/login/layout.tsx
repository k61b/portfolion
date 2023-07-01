export const metadata = {
  title: 'Login',
  description: 'Portfolion login page',
}

export default function DashboardLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <section className="flex flex-col h-screen">
      <main className="flex flex-col flex-grow">{children}</main>
    </section>
  )
}
