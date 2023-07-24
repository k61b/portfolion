import Popup from '../addBookmark/Popup'
import List from './list'

export default function Table() {
  return (
    <div>
      <div className="p-4 bg-slate-900 rounded-b shadow-xl flex flex-row items-center justify-between">
        <h4 className="scroll-m-20 text-xl font-semibold tracking-tight text-green-400">
          Dashboard
        </h4>
        <h3 className="text-white scroll-m-20 text-2xl font-semibold tracking-tight">
          PORTFOLION
        </h3>
        <p className="text-white">
          <Popup />
        </p>
      </div>
      <List />
    </div>
  )
}
