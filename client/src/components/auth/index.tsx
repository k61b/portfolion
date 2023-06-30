'use client'
import { POST } from '@app/(auth)/route'
import { useFormik } from 'formik'
import * as Yup from 'yup'

export default function AuthForm() {
  const formik = useFormik({
    initialValues: {
      username: '',
      password: '',
    },
    validationSchema: Yup.object({
      username: Yup.string().required('Required'),
      password: Yup.string().required('Required'),
    }),
    onSubmit: async (body) => {
      await POST('/api/session', body)
      console.log(body)
    },
  })

  return (
    <form
      className="flex flex-col items-center justify-center"
      onSubmit={formik.handleSubmit}
    >
      <input
        id="username"
        name="username"
        type="text"
        placeholder="Username"
        onChange={formik.handleChange}
        value={formik.values.username}
        className="border border-gray-300 rounded-md p-2 mb-2 text-black"
      />
      {formik.errors.username ? <div>{formik.errors.username}</div> : null}
      <input
        id="password"
        name="password"
        type="password"
        placeholder="Password"
        onChange={formik.handleChange}
        value={formik.values.password}
        className="border border-gray-300 rounded-md p-2 mb-2 text-black"
      />
      {formik.errors.password ? <div>{formik.errors.password}</div> : null}
      <button type="submit" className="bg-blue-500 rounded-md p-2 text-black">
        Login
      </button>
    </form>
  )
}
