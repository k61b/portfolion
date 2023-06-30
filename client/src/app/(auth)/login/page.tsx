import AuthForm from "@components/auth";

export default function Login() {
  return (
    <div className="flex flex-col items-center justify-center h-screen">
      <h1 className="text-3xl font-bold mb-4">Login Page</h1>
      <AuthForm />
    </div>
  )
}
