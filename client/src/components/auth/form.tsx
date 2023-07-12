/* eslint-disable @typescript-eslint/no-misused-promises */
import { postFetcher } from '@/service'
import { AuthContext } from '@/context/auth'
import { useContext, useState } from 'react'
import { Link } from 'react-router-dom'
import { useNavigate } from 'react-router-dom'
import * as z from 'zod'
import { zodResolver } from '@hookform/resolvers/zod'
import { useForm } from 'react-hook-form'
import { Button } from '@/components/ui/button'
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from '@/components/ui/form'
import { Input } from '@/components/ui/input'
import { useTranslation } from 'react-i18next';

const schema = z.object({
  username: z
    .string()
    .min(3, {
      message: 'Username must be at least 3 characters long',
    })
    .max(50),
  password: z
    .string()
    .min(4, {
      message: 'Password must be at least 8 characters long',
    })
    .max(50),
})

export function AuthForm() {
  const navigate = useNavigate()
  const { setUser } = useContext(AuthContext)
  const [error, setError] = useState(null)
  const { t } = useTranslation();

  const form = useForm<z.infer<typeof schema>>({
    resolver: zodResolver(schema),
    defaultValues: {
      username: '',
      password: '',
    },
  })

  function onSubmit(data: z.infer<typeof schema>) {
    postFetcher('/api/v1/session', data)
      .then((data) => {
        if (data.error) {
          setError(data.error)
        } else {
          setUser(data)
          navigate('/dashboard')
        }
      })
      .catch((err) => {
        setError(err.message)
      })
  }

  if (error) {
    return (
      <div className="flex flex-col items-center justify-center h-screen">
        <div className="flex flex-col items-center justify-center space-y-4">
          <h1 className="text-4xl font-bold">Error</h1>
          <p className="text-lg">{t("login.error.title")}</p>
          <Link to="/login" className="text-slate-700">
            {t("login.error.back-to-login")}
          </Link>
        </div>
      </div>
    )
  }

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8 w-96 border p-4 rounded-lg shadow-slate-600">
        <FormField
          control={form.control}
          name="username"
          render={({ field }) => (
            <FormItem>
              <FormLabel>{t("login.username")}</FormLabel>
              <FormControl>
                <Input placeholder="username" {...field} />
              </FormControl>
              <FormDescription>
                {t("login.your", {text: t("login.username")})}
              </FormDescription>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="password"
          render={({ field }) => (
            <FormItem>
              <FormLabel>{t("login.password")}</FormLabel>
              <FormControl>
                <Input placeholder="password" {...field} type='password' />
              </FormControl>
              <FormDescription>{t("login.your", {text: t("login.password")})}</FormDescription>
              <FormMessage />
            </FormItem>
          )}
        />

        <Button type="submit">{t('login.title')}</Button>
      </form>
    </Form>
  )
}
