/* eslint-disable @typescript-eslint/no-unsafe-return */
import axios, { AxiosRequestConfig } from 'axios'

export const axiosInstance = axios.create({
  baseURL: import.meta.env.VITE_PUBLIC_API_URL,
  withCredentials: true,
})

axiosInstance.interceptors.response.use(
  function (response) {
    return response
  },
  function (error) {
    return Promise.reject(error.response.data)
  }
)

export const fetcher = async (url: any, config?: AxiosRequestConfig) => {
  const res = await axiosInstance.get(url, config)
  return res.data
}

export const postFetcher = async (
  url: any,
  data: any,
  config?: AxiosRequestConfig
) => {
  const res = await axiosInstance.post(url, data, config)
  return res.data
}
