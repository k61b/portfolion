import { AuthForm } from '@/components/auth/form'

export default function Login() {
  return (
    <div
      id="root"
      className="flex flex-col items-center justify-center min-h-screen py-4 bg-slate-50 dark:bg-slate-800 text-slate-950 dark:text-slate-50"
    >
      <AuthForm />
    </div>
  )
}
