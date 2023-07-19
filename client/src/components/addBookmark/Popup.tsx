import { useState } from 'react'
import Modal from './Modal'

export default function Popup() {
  const [showModal, setShowModal] = useState(false)
  return (
    <div>
      <div>
        <div>
          <button
            onClick={() => setShowModal(true)}
            className="bg-white text-slate-900 px-3 py-2 rounded hover:scale-95 transition font-bold"
          >
            +
          </button>
        </div>
      </div>
      <Modal visible={showModal} />
    </div>
  )
}
