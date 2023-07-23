import { fetcher } from '@/service'
import useSWR from 'swr'

export const useBookmark = () => {
  const { data, error, isLoading, mutate } = useSWR(`/api/v1/bookmarks`, fetcher)

  return {
    error,
    isLoading,
    bookmarks: data,
    mutateBookmark: mutate,
  }
}
