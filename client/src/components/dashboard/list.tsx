import { useBookmark } from '@/hooks/data/useBookmark'
import { cn } from '@/lib/utils'
import { useTranslation } from 'react-i18next'

export default function List() {
  const { bookmarks } = useBookmark()
  const { t } = useTranslation()

  return (
    <div>
      {bookmarks ? (
        <div className="flex flex-col justify-center mt-3 p-9">
          <div className="bg-slate-900 text-white rounded">
            <ul className="grid grid-cols-5 gap-4 p-4 font-medium">
              <li className="text-center">{t('dashboard.table.symbol')}</li>
              <li className="text-center">
                {t('dashboard.table.added_price')}
              </li>
              <li className="text-center">{t('dashboard.table.pieces')}</li>
              <li className="text-center">
                {t('dashboard.table.current_price')}
              </li>
              <li className="text-center">
                {t('dashboard.table.profit_and_loss')}
              </li>
            </ul>
          </div>
          {bookmarks?.map((bookmark: any, index: number) => (
            <ul
              key={index}
              className="grid grid-cols-5 gap-4 items-center justify-between bg-white p-4 rounded shadow-md mt-4"
            >
              <li className="text-center font-semibold">{bookmark.symbol}</li>
              <li className="text-center">{bookmark.added_price}</li>
              <li className="text-center">{bookmark.pieces}</li>
              <li className="text-center">{bookmark.current_price}</li>
              <li
                className={cn('text-center text-green-500 font-medium', {
                  'text-red-500': bookmark.profit_and_loss < 0,
                })}
              >
                {bookmark.profit_and_loss.toString().length > 7 ? (
                  <>
                    <span className="hidden sm:inline">
                      {bookmark.profit_and_loss.toString().substring(0, 7)}
                    </span>
                    <span className="inline sm:hidden">
                      {bookmark.profit_and_loss.toString().substring(0, 4)}
                    </span>
                    ...
                  </>
                ) : (
                  bookmark.profit_and_loss
                )}
              </li>
            </ul>
          ))}
        </div>
      ) : (
        <div className="flex flex-col justify-center items-center h-screen text-slate-900">
          <h1 className="text-xl md:text-2xl p-4">
            Add a stock to your portfolio by clicking the plus button in the top
          </h1>
        </div>
      )}
    </div>
  )
}
