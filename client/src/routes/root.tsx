import Logo from '@/components/logo/intex'

export default function Root() {
  return (
    <div
      id="root"
      className="flex flex-col items-center min-h-screen py-4 bg-slate-50 dark:bg-slate-800 text-slate-950 dark:text-slate-50"
    >
      <Logo />
    </div>
  )
}
