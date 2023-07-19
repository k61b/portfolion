import BookmarkForm from './BookmarkForm'

export default function Modal({ visible = false }) {
  if (!visible) return null
  return (
    <div className="fixed inset-0 bg-black bg-opacity-30 backdrop-blur-sm flex justify-center items-center">
      <div className="bg-white rounded p-10">
        <h3
          className="scroll-m-20 text-2xl font-semibold tracking-tight text-center text-slate-900 mb-4
        border-b pb-2 transition-colors first:mt-0"
        >
          Add Bookmark
        </h3>
        <div>
          <BookmarkForm />
        </div>
        <div className="flex flex-row justify-around items-center"></div>
      </div>
    </div>
  )
}
