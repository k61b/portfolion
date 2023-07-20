import { Avatar, AvatarImage } from '@radix-ui/react-avatar'
import { Button } from '../ui/button'
import { useContext } from 'react'
import { AuthContext } from '@/context/auth'
import { fetcher } from '@/service'
import { useNavigate } from 'react-router-dom'
import { useTranslation } from 'react-i18next'
import { cn } from '@/lib/utils'

export default function Sidebar() {
  const { user } = useContext(AuthContext)
  const { setUser } = useContext(AuthContext)
  const { t } = useTranslation()
  const navigate = useNavigate()

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
    <div className="flex flex-row justify-around items-center md:items-stretch md:flex-col p-4 bg-slate-50 rounded shadow-xl h-max md:h-screen md:w-auto">
      <Avatar>
        <AvatarImage className="w-36" src={user?.avatar} alt="avatar" />
      </Avatar>
      <div className="mt-4 flex flex-col justify-center">
        <h1 className="text-2xl font-semibold text-slate-900 mb-3">
          {user?.username}
        </h1>

        <span className="text-slate-700 mb-1">
          {t('dashboard.user.total', { value: user?.value })}
        </span>

        <span
          className={cn('text-green-700', {
            'text-red-700': user && user.profit_and_loss < 0,
          })}
        >
          {t('dashboard.user.profit_and_loss', {
            value: user?.profit_and_loss,
          })}
        </span>
      </div>
      <Button
        onClick={handleLogout}
        className="bg-slate-900 mt-auto w-max md:w-full"
      >
        {t('dashboard.logout')}
      </Button>
    </div>
  )
}
