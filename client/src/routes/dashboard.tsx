import { AuthContext } from '@/context/auth'
import { useContext } from 'react'

export default function Dashboard() {
  const { user } = useContext(AuthContext)
  return (
    <div id="dashboard">
      <h1>Welcome {user?.username}</h1>
    </div>
  )
}
