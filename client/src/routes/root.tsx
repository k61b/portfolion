import Logo from '@/components/logo/intex'
import { Link } from 'react-router-dom'

export default function Root() {
  return (
    <div
      id="root"
      className="flex flex-col items-center min-h-screen py-4 bg-slate-50 dark:bg-slate-800 text-slate-950 dark:text-slate-50"
    >
      <Logo />
      <Link
        to={'/login'}
        className="italic underline underline-offset-1 text-slate-800"
      >
        go to login
      </Link>
    </div>
  )
}
