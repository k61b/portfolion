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
    <div className="flex flex-row justify-around items-center md:items-stretch md:flex-col p-4 bg-slate-50 rounded shadow-xl h-max md:h-screen ">
      <Avatar>
        <AvatarImage
          className="w-36"
          src={`https://api.dicebear.com/6.x/big-smile/svg?seed=${user?.avatar}&radius=50&backgroundType=gradientLinear&accessories=catEars,glasses,mustache,sailormoonCrown,sleepMask,sunglasses&eyes=angry,confused,normal,sad,sleepy,starstruck,winking,cheery&hairColor=220f00,238d80,3a1a00,605de4,71472d,d56c0c&skinColor=8c5a2b,a47539,c99c62,e2ba87,efcc9f,643d19,f5d7b1&backgroundColor=ffd5dc,d1d4f9,c0aede,b6e3f4`}
          alt='avatar'
        />
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
