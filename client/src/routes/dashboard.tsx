import { Button } from '@/components/ui/button'
import { AuthContext } from '@/context/auth'
import { fetcher } from '@/service'
import { useContext } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import { useTranslation } from 'react-i18next'

export default function Dashboard() {
  const { user } = useContext(AuthContext)
  const { setUser } = useContext(AuthContext)
  const navigate = useNavigate()
  const { t } = useTranslation()

  const handleLogout = () => {
    fetcher('/api/v1/logout')
      .then((data) => {
        if (data.error) {
          console.error(data.error)
        } else {
          setUser(null)
          navigate('/')
        }
      })
      .catch((err) => {
        console.error(err.message)
      })
  }

  return (
    <div className="bg-slate-50">
      {user ? (
        <div className="flex flex-col justify-center items-center h-screen">
          <h1>{t('dashboard.title')}</h1>
          <p>{t('dashboard.welcome', { name: user.username })}</p>
          <Button onClick={handleLogout}>{t('dashboard.logout')}</Button>
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
