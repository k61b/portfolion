'use client'
import { fetcher } from '@utils/fetch'
import useSWR from 'swr'

export default function Dashboard() {
  const { data: user } = useSWR('/api/auth', fetcher)

  return (
    <div className="flex flex-col flex-grow">
      <h1 className="text-3xl font-bold">Dashboard</h1>
      <p className="text-xl">Welcome {user?.username}</p>
    </div>
  )
}
