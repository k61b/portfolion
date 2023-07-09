import { fetcher } from '@/service'
import { IUser } from '@/types/user'
import useSWR from 'swr'

export const useUser = () => {
  const { data, error, isLoading, mutate } = useSWR<IUser>(`/api/v1/auth`, fetcher)

  return {
    error,
    isLoading,
    user: data,
    mutateUser: mutate,
  }
}
