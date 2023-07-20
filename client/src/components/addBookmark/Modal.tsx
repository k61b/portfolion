import { useTranslation } from 'react-i18next'
import BookmarkForm from './BookmarkForm'

export default function Modal({ visible = false }) {
  const { t } = useTranslation()

  if (!visible) return null
  return (
    <div className="fixed inset-0 bg-black bg-opacity-30 backdrop-blur-sm flex justify-center items-center">
      <div className="bg-white rounded p-10">
        <h3
          className="scroll-m-20 text-2xl font-semibold tracking-tight text-center text-slate-900 mb-4
        border-b pb-2 transition-colors first:mt-0"
        >
          {t('popup.title')}
        </h3>
        <div>
          <BookmarkForm />
        </div>
      </div>
    </div>
  )
}
