import useSWR from 'swr'
import { fetcher } from '@utils/fetch'

export interface User {
  username: string
  bookmarks: {
    symbol: string
    added_price: number
    pices: number
  }
}

export const useAuth = () => {
  const { data, error, isLoading, mutate } = useSWR<User>(`/api/auth`, fetcher)

  return {
    error,
    isLoading,
    user: data,
    mutateMe: mutate,
  }
}
