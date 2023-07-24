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
            className="bg-green-500 hover:bg-green-400 text-white font-bold py-2 px-4 rounded"
          >
            +
          </button>
        </div>
      </div>
      <Modal visible={showModal} />
    </div>
  )
}
