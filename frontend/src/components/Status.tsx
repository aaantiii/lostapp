import '@styles/components/Status.scss'
import { OrderStatus as OS, RecruitmentRequestStatus } from '@api/types/models'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faDotCircle } from '@fortawesome/free-solid-svg-icons'
import { forwardRef } from 'react'
import { TFunction } from 'i18next'

type StatusProps = {
  status: OS | RecruitmentRequestStatus
  tFunc: TFunction
  tPrefix?: string
  onClick?: () => void
  disabled?: boolean
}

const Status = forwardRef<HTMLButtonElement, StatusProps>(({ status, onClick, disabled, tFunc, tPrefix }, ref) => {
  if (tPrefix && !tPrefix.endsWith('.')) tPrefix += '.'
  return (
    <button className={`Status ${status}`} onClick={onClick} ref={ref} disabled={disabled} aria-disabled={disabled || !onClick}>
      <FontAwesomeIcon icon={faDotCircle} className="icon" />
      <span className="text">{tFunc((tPrefix ?? '') + status)}</span>
    </button>
  )
})

export default Status
