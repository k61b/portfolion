import { createContext, ReactNode, useState } from 'react'
import { useUser } from '@/hooks/data/useUser'
import { IUser } from '@/types/user'

interface IAuthContext {
  user: IUser | null
  setUser: (user: IUser | null) => void
}

const initialAuthContext: IAuthContext = {
  user: null,
  setUser: () => {},
}

export const AuthContext = createContext<IAuthContext>(initialAuthContext)

export const AuthProvider = ({ children }: { children: ReactNode }) => {
  const { user } = useUser()
  const [authUser, setAuthUser] = useState<IUser | null>(
    initialAuthContext.user
  )

  if (user && !authUser) {
    setAuthUser(user)
  }

  return (
    <AuthContext.Provider
      value={{
        user: authUser,
        setUser: setAuthUser,
      }}
    >
      {children}
    </AuthContext.Provider>
  )
}
