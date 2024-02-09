import '@styles/components/AlertDialog.scss'
import * as ad from '@radix-ui/react-alert-dialog'
import { ReactNode, useState } from 'react'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faXmark } from '@fortawesome/free-solid-svg-icons'

type AlertDialogProps = {
  trigger: ReactNode
  title: string
  description?: string
  confirm: ReactNode
  cancel?: ReactNode
  children?: ReactNode
}

export default function AlertDialog({ title, trigger, description, confirm, cancel, children }: AlertDialogProps) {
  const [open, setOpen] = useState(false)

  return (
    <ad.Root open={open}>
      <ad.Trigger asChild onClick={() => setOpen(true)}>
        {trigger}
      </ad.Trigger>
      <ad.Portal>
        <ad.Overlay className="AlertDialogOverlay" />
        <ad.Content className="AlertDialogContent">
          <div className="header">
            <span>{title}</span>
            <ad.Cancel asChild className="close" onClick={() => setOpen(false)}>
              <button aria-label="Close">
                <FontAwesomeIcon icon={faXmark} />
              </button>
            </ad.Cancel>
          </div>
          <div className="scroll-wrapper">
            <ad.Title className="title">{title}</ad.Title>
            {description && <ad.Description className="description">{description}</ad.Description>}
            <div className="content">{children}</div>
            <div className="buttons">
              {cancel && (
                <ad.Cancel asChild onClick={() => setOpen(false)}>
                  {cancel}
                </ad.Cancel>
              )}
              <ad.Action asChild onClick={() => setOpen(false)}>
                {confirm}
              </ad.Action>
            </div>
          </div>
        </ad.Content>
      </ad.Portal>
    </ad.Root>
  )
}
