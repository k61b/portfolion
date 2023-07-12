import { AuthForm } from '@/components/auth/form'
import { useTranslation } from 'react-i18next';

export default function Login() {
  const { t } = useTranslation();
  return (
    <div
      id="root"
      className="flex flex-col items-center justify-center min-h-screen py-4 bg-slate-50 dark:bg-slate-800 text-slate-950 dark:text-slate-50"
    >
      <h1 className='scroll-m-20 text-4xl font-extrabold tracking-tight lg:text-4xl p-4'>{t('login.title')}</h1>
      <AuthForm />
    </div>
  )
}
