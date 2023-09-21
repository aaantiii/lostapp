import '@styles/components/Dialog.scss'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import Button from './Button'
import * as d from '@radix-ui/react-dialog'
import { faXmark } from '@fortawesome/free-solid-svg-icons'
import { useCallback, useEffect, useState } from 'react'

type ConfirmDialogProps = {
  title: string
  description: string
  confirmText: string
  cancelText: string
  triggerButtonColor?: 'red'
  onOpenChange?: (open: boolean) => void
  onConfirm?: () => void
}

export default function ConfirmDialog(props: ConfirmDialogProps) {
  const [open, setOpen] = useState(false)

  useEffect(() => {
    window.addEventListener('keydown', handleKeyDown)
    return () => window.removeEventListener('keydown', handleKeyDown)
  }, [])

  const handleKeyDown = useCallback((e: KeyboardEvent) => {
    if (e.key === 'Escape') setOpen(false)
  }, [])

  function handleConfirm() {
    props.onConfirm?.()
    setOpen(false)
  }

  return (
    <d.Root open={open} onOpenChange={props.onOpenChange}>
      <d.Trigger asChild>
        <Button className={props.triggerButtonColor} onClick={() => setOpen(true)}>
          {props.title}
        </Button>
      </d.Trigger>
      <d.Portal>
        <d.Overlay className="DialogOverlay" />
        <d.Content className="DialogContent" onPointerDownOutside={() => setOpen(false)}>
          <d.Title className="title">{props.title}</d.Title>
          <d.Description className="description">{props.description}</d.Description>
          <div className="buttons">
            <Button className="red" onClick={() => setOpen(false)}>
              {props.cancelText}
            </Button>
            <Button onClick={handleConfirm}>{props.confirmText}</Button>
          </div>
          <d.Close asChild className="close">
            <button aria-label="Schließen" onClick={() => setOpen(false)}>
              <FontAwesomeIcon icon={faXmark} />
            </button>
          </d.Close>
        </d.Content>
      </d.Portal>
    </d.Root>
  )
}
