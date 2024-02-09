import '@styles/components/Dialog.scss'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import * as d from '@radix-ui/react-dialog'
import { faXmark } from '@fortawesome/free-solid-svg-icons'
import { FormProps } from './Form'
import { Suspense, useState } from 'react'
import { LazyForm } from './LazyComponents'

type DialogProps = FormProps & {
  title: string
  trigger: React.ReactNode
  description?: string
  onOpenChange?: (open: boolean) => void
  closeOnEscape?: boolean
}

export default function Dialog({ title, description, trigger, closeOnEscape, onOpenChange, onSubmit, ...formProps }: DialogProps) {
  const [open, setOpen] = useState(false)

  async function handleSubmit(data: FormData) {
    const res = onSubmit?.(data)

    if (res instanceof Promise) {
      const success = await res
      if (success) setOpen(false)
    } else if (res === true) return setOpen(false)
  }

  return (
    <d.Root open={open} onOpenChange={onOpenChange}>
      <d.Trigger asChild onClick={() => setOpen(true)}>
        {trigger}
      </d.Trigger>
      <d.Portal>
        <d.Overlay className="DialogOverlay" />
        <d.Content
          className="DialogContent"
          onPointerDownOutside={() => setOpen(false)}
          onEscapeKeyDown={closeOnEscape ? undefined : (e) => e.preventDefault()}
        >
          <div className="wrapper">
            <div className="header">
              <span>{title}</span>
              <d.Close asChild className="close" onClick={() => setOpen(false)}>
                <button aria-label="Close">
                  <FontAwesomeIcon icon={faXmark} />
                </button>
              </d.Close>
            </div>
            <div className="scroll-wrapper">
              <d.Title className="title">{title}</d.Title>
              {description && <d.Description className="description">{description}</d.Description>}
              <Suspense>
                <LazyForm {...formProps} onSubmit={handleSubmit} />
              </Suspense>
            </div>
          </div>
        </d.Content>
      </d.Portal>
    </d.Root>
  )
}
