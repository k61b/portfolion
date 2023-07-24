import Logo from '@/components/logo/intex'
import { Link } from 'react-router-dom'
import styles from './root.module.css'

export default function Root() {
  return (
    <div className={styles.container}>
      <div className=" flex flex-col justify-center items-center h-screen text-white">
        <Logo />
        <h1 className="text-4xl font-semibold text-center">
          Welcome to <span className="text-green-400">
            Portfolion
          </span>
        </h1>
        <p className="text-center mt-4">
          <span className="text-green-400">Portfolion</span> is a simple web app
          that allows you to keep track of your stock portfolio.
        </p>
        <div className="mt-4">
          <Link
            to={'/login'}
            className="bg-green-500 hover:bg-green-400 text-white font-bold py-2 px-4 rounded"
          >
            Get Started
          </Link>
        </div>
      </div>
    </div>
  )
}
