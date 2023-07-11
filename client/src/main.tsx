import React from 'react'
import './index.css'
import ReactDOM from 'react-dom/client'
import Root from './routes/root'
import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import ErrorPage from './pages/error-page'
import Login from './routes/login'
import Dashboard from './routes/dashboard'
import { AuthProvider } from './context/auth'
import '@/lib/i18n'

const router = createBrowserRouter([
  {
    path: '/',
    element: <Root />,
    errorElement: <ErrorPage />,
  },
  {
    path: '/login',
    element: <Login />,
    errorElement: <ErrorPage />,
  },
  {
    path: '/dashboard',
    element: (
      <AuthProvider>
        <Dashboard />
      </AuthProvider>
    ),
    errorElement: <ErrorPage />,
  },
])

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
)
