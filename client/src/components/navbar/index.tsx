'use client'
import { useAuth } from '@hooks/data/useAuth'
import { fetcher } from '@utils/fetch'
import Link from 'next/link'

export default function Navbar() {
  const { user } = useAuth()

  const handleLogout = async () => {
    await fetcher('/logout')

    window.location.href = '/login'
  }

  return (
    <nav className="flex items-center justify-between flex-wrap bg-gray-800 p-6 m-6 rounded-3xl">
      <div className="flex items-center flex-shrink-0 text-white mr-6">
        <Link
          href="/dashboard"
          className="font-semibold text-xl tracking-tight"
        >
          Portfolion
        </Link>
      </div>

      <div className="flex flex-col items-center flex-shrink-0 text-white">
        {user && user.profit_and_loss < 0 ? (
          <span className="font-semibold text-sm tracking-tight text-red-500">
            {user?.profit_and_loss}
          </span>
        ) : (
          <span className="font-semibold text-sm tracking-tight text-green-500">
            {user?.profit_and_loss}
          </span>
        )}

        {user && (
          <span className="font-semibold text-2xl tracking-tight">
            {user.value}
          </span>
        )}
      </div>

      <div className="flex items-center flex-shrink-0 text-white mr-6">
        {user ? (
          <div className="flex flex-col items-center gap-2">
            <span className="font-semibold text-xl tracking-tight">
              {user.username}
            </span>
            <button
              className="bg-gray-700 hover:bg-gray-600 text-white font-bold py-2 px-4 rounded"
              onClick={handleLogout}
            >
              Logout
            </button>
          </div>
        ) : (
          <Link href="/login" className="font-semibold text-xl tracking-tight">
            Login
          </Link>
        )}
      </div>
    </nav>
  )
}
