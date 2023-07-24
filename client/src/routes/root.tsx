import Logo from '@/components/logo/intex'
import { Link } from 'react-router-dom'
import styles from './root.module.css'
import { useTranslation } from 'react-i18next'

export default function Root() {
  const { t } = useTranslation()
  return (
    <div className={styles.container}>
      <div className=" flex flex-col justify-center items-center h-screen text-white">
        <Logo />
        <h1 className="text-4xl font-semibold text-center text-green-400">
          {t('root.welcome', { text: t('root.title') })}
        </h1>
        <p className="text-center mt-4">
          <span className="text-green-400">{t('root.title')},</span>{' '}
          {t('root.description')}
        </p>
        <div className="mt-4">
          <Link
            to={'/login'}
            className="bg-green-500 hover:bg-green-400 text-white font-bold py-2 px-4 rounded"
          >
            {t('root.link')}
          </Link>
        </div>
      </div>
    </div>
  )
}
