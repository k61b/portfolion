import { Button } from '@/components/ui/button'
import { AuthContext } from '@/context/auth'
import { fetcher } from '@/service'
import { useContext } from 'react'
import { Link, useNavigate } from 'react-router-dom'

export default function Dashboard() {
  const { user } = useContext(AuthContext)
  const { setUser } = useContext(AuthContext)
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
    <div className="bg-slate-50">
      {user ? (
        <div className="flex flex-col justify-center items-center h-screen">
          <h1>Dashboard</h1>
          <p>Welcome {user.username}</p>
          <Button onClick={handleLogout}>Logout</Button>
        </div>
      ) : (
        <div className="flex flex-col justify-center items-center h-screen">
          <h1>you must be logged in to see the Dashboard</h1>

          <Link
            to={'/login'}
            className="italic underline underline-offset-1 text-slate-800"
          >
            go to login page
          </Link>
        </div>
      )}
    </div>
  )
}
