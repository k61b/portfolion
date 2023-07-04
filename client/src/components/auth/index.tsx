"use client";
import { postFetcher } from "@utils/fetch";
import { useFormik } from "formik";
import { useRouter } from "next/navigation";
import * as Yup from "yup";
import { Inter } from "next/font/google";

const inter = Inter({
  subsets: ["latin"],
  display: "swap",
});

export default function AuthForm() {
  const router = useRouter();
  const formik = useFormik({
    initialValues: {
      username: "",
      password: "",
    },
    validationSchema: Yup.object({
      username: Yup.string().required("Required"),
      password: Yup.string().required("Required"),
    }),
    onSubmit: async (body) => {
      await postFetcher(`/api/session`, body);
      router.push("/dashboard");
    },
  });

  return (
    <div className="bg-white dark:bg-slate-800 shadow-slate-600 shadow-2xl max-h-screen h-max max-w-max p-8 rounded-lg ">
      <div className={inter.className}>
        <h1 className="text-3xl font-bold mb-4 dark:text-white">
          Welcome to <span className="text-blue-500">Portfolion</span>
        </h1>
        <form onSubmit={formik.handleSubmit}>
          <div className="md:block grid grid-cols-2 gap-4 mb-4">
            <div className="flex flex-col">
              <label
                htmlFor="username"
                className="block text-gray-700 dark:text-gray-200 text-sm font-bold mt-3 dark:text-white"
              >
                Username
              </label>
              <input
                id="username"
                name="username"
                type="text"
                className="shadow border rounded w-full py-2 px-3 text-gray-700 dark:text-gray-200 leading-tight focus:outline-none focus:shadow-outline mt-3"
                placeholder="johndoe"
                onChange={formik.handleChange}
                value={formik.values.username}
              />

              {formik.touched.username && formik.errors.username ? (
                <div className="text-red-500 text-xs italic mt-2">
                  {formik.errors.username}
                </div>
              ) : null}
            </div>
            <div className="flex flex-col">
              <label
                htmlFor="password"
                className="block text-gray-700 dark:text-gray-200 text-sm font-bold mt-3"
              >
                Password
              </label>
              <input
                id="password"
                type="password"
                name="password"
                onChange={formik.handleChange}
                value={formik.values.password}
                className="shadow border rounded w-full py-2 px-3 text-gray-700 dark:text-gray-200 leading-tight focus:outline-none focus:shadow-outline mt-3"
                placeholder="*********"
              />

              {formik.touched.password && formik.errors.password ? (
                <div className="text-red-500 text-xs italic mt-2">
                  {formik.errors.password}
                </div>
              ) : null}
            </div>
          </div>
          <div className="flex items-center justify-between mt-2 sm:flex-col">
            <label className="block text-gray-500 font-bold">
              <input className="mr-2" type="checkbox" name="remember"/>
              <span>Remember Me</span>
            </label>

            <a
              className="inline-block align-baseline font-bold text-sm text-blue-500 hover:text-blue-800"
              href="#"
            >
              Forgot Password?
            </a>
          </div>
          <button
            type="submit"
            className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline w-full mt-4"
          >
            Sign In
          </button>
        </form>
        <div className="flex items-center mt-12">
          <hr className="flex-grow border-t border-gray-300" />
          <p className="mx-4 text-center text-gray-500 text-lg">or</p>
          <hr className="flex-grow border-t border-gray-300" />
        </div>
        <div className="flex flex-col items-center justify-center mt-4">
          <button
            type="button"
            className="border border-red-600 text-red-600 font-normal py-2 px-4 rounded focus:outline-none focus:shadow-outline w-full mt-6 hover:bg-red-600 hover:text-white"
          >
            Sign In with Google
          </button>
        </div>
        <div className="flex flex-col items-center justify-center mt-4">
          <p
            className="
          text-gray-500 text-lg
          "
          >
            Don't have an account ?{" "}
            <a href="#" className="text-blue-500 hover:text-blue-800">
              Create Account
            </a>
          </p>
        </div>
      </div>
    </div>
  );
}
