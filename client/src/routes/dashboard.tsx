import { AuthContext } from '@/context/auth'
import { useContext } from 'react'
import { useTranslation } from 'react-i18next'
import Sidebar from '@/components/dashboard/sidebar'
import Table from '@/components/dashboard/table'
import { Link } from 'react-router-dom'

export default function Dashboard() {
  const { user } = useContext(AuthContext)
  const { t } = useTranslation()
  console.log(user?.avatar)
  return (
    <div className="bg-slate-50">
      {user ? (
        <div className="grid md:grid-cols-[auto,1fr]">
          <div className="mb-4 md:mb-0">
            <Sidebar />
          </div>
          <div className="w-full ">
            <Table />
          </div>
        </div>
      ) : (
        <div className="flex flex-col justify-center items-center h-screen">
          <h1>{t('dashboard.error.title')}</h1>

          <Link
            to={'/login'}
            className="italic underline underline-offset-1 text-slate-800"
          >
            {t('dashboard.error.back-to-login')}
          </Link>
        </div>
      )}
    </div>
  )
}
