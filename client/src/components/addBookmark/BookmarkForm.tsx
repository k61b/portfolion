import { postFetcher } from '@/service'
import { zodResolver } from '@hookform/resolvers/zod'
import { useState } from 'react'
import { useForm } from 'react-hook-form'
import * as z from 'zod'
import { Button } from '@/components/ui/button'
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from '@/components/ui/form'
import { Input } from '@/components/ui/input'
import { useTranslation } from 'react-i18next'

const formSchema = z.object({
  added_price: z.string().min(1, {
    message: 'Must have at least 1',
  }),
  pieces: z.string().min(1, {
    message: 'Must have at least 1',
  }),
  symbol: z.string().min(1, {
    message: 'Symbol must be at least 1 characters.',
  }),
})

export default function BookmarkForm() {
  const [error, setError] = useState(null)
  const { t } = useTranslation()

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      added_price: '',
      pieces: '',
      symbol: '',
    },
  })

  function onSubmit(values: z.infer<typeof formSchema>) {
    const postData = {
      added_price: Number(values.added_price),
      pieces: Number(values.pieces),
      symbol: values.symbol,
    }
    postFetcher('/api/v1/bookmarks', postData).then((data) => {
      if (data.error) {
        setError(data.error)
      } else {
        window.location.reload()
      }
    })
  }

  if (error) {
    return <div>{error}</div>
  }

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)}>
        <FormField
          control={form.control}
          name="symbol"
          render={({ field }) => (
            <FormItem className='mb-8'>
              <FormLabel className='text-slate-900'>{t('popup.symbol.title')}</FormLabel>
              <FormControl>
                <Input
                  placeholder={t('popup.symbol.placeholder')}
                  {...field}
                  className="text-slate-900"
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <div className="md:flex mb-8 space-x-4">
          <FormField
            control={form.control}
            name="added_price"
            render={({ field }) => (
              <FormItem>
                <FormLabel className='text-slate-900'>{t('popup.added_price.title')}</FormLabel>
                <FormControl>
                  <Input
                    placeholder={t('popup.added_price.placeholder')}
                    {...field}
                    className="text-slate-900"
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="pieces"
            render={({ field }) => (
              <FormItem>
                <FormLabel className='text-slate-900'>{t('popup.pieces.title')}</FormLabel>
                <FormControl>
                  <Input
                    placeholder={t('popup.pieces.placeholder')}
                    {...field}
                    className="text-slate-900"
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
        </div>

        <div className="flex flex-row justify-around items-center">
          <Button className="bg-green-700 hover:bg-green-800" type="submit">
            {t('popup.button.add')}
          </Button>
          <Button
            className="bg-red-700 hover:bg-red-800"
            onClick={() => {
              window.location.reload()
            }}
          >
            {t('popup.button.close')}
          </Button>
        </div>
      </form>
    </Form>
  )
}
