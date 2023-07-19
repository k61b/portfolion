import { postFetcher } from '@/service'
import { zodResolver } from '@hookform/resolvers/zod'
import { useState } from 'react'
import { useForm } from 'react-hook-form'
import * as z from 'zod'
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
            <FormItem>
              <FormLabel>Symbol</FormLabel>
              <FormControl>
                <Input
                  placeholder="search symbol"
                  {...field}
                  className="text-slate-900"
                />
              </FormControl>
              <FormDescription>
                Symbol of the investment you will add
              </FormDescription>
              <FormMessage />
            </FormItem>
          )}
        />
        <div className="md:flex mb-8">
          <FormField
            control={form.control}
            name="added_price"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Added Price</FormLabel>
                <FormControl>
                  <Input
                    placeholder="added price"
                    {...field}
                    className="text-slate-900"
                  />
                </FormControl>
                <FormDescription>Price at time of purchase</FormDescription>
                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="pieces"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Pieces</FormLabel>
                <FormControl>
                  <Input
                    placeholder="pieces"
                    {...field}
                    className="text-slate-900"
                  />
                </FormControl>
                <FormDescription>Quantity purchased</FormDescription>
                <FormMessage />
              </FormItem>
            )}
          />
        </div>

        <div className="flex flex-row justify-around items-center">
          <Button className="bg-green-700 hover:bg-green-800" type="submit">
            Add
          </Button>
          <Button
            className="bg-red-700 hover:bg-red-800"
            onClick={() => {
              window.location.reload()
            }}
          >
            Close
          </Button>
        </div>
      </form>
    </Form>
  )
}
