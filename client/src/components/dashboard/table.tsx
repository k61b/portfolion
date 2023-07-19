import Popup from "../addBookmark/Popup";

export default function Table() {
  return (
    <div className="p-4 bg-slate-900 rounded shadow-xl flex flex-row items-center justify-between">
      <p className="text-white">dashboard</p>
      <h3 className="text-white">PORTFOLION</h3>
      <p className="text-white">
        <Popup />
      </p>
    </div>
  )
}
