import '@styles/components/Dialog.scss'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import Button from './Button'
import * as d from '@radix-ui/react-dialog'
import { faXmark } from '@fortawesome/free-solid-svg-icons'
import Form, { FormProps } from './Form'
import { useState } from 'react'

interface FormDialogProps extends FormProps {
  title: string
  description: string
  onOpenChange?: (open: boolean) => void
}

export default function FormDialog(props: FormDialogProps) {
  const [open, setOpen] = useState(false)

  function handleSubmit(data: any) {
    props.onSubmit?.(data)
    setOpen(false)
  }

  return (
    <d.Root open={open} onOpenChange={props.onOpenChange}>
      <d.Trigger asChild>
        <Button onClick={() => setOpen(true)}>{props.title}</Button>
      </d.Trigger>
      <d.Portal>
        <d.Overlay className="DialogOverlay" />
        <d.Content className="DialogContent" onPointerDownOutside={() => setOpen(false)}>
          <d.Title className="title">{props.title}</d.Title>
          <d.Description className="description">{props.description}</d.Description>
          {<Form submitText={props.submitText} onSubmit={handleSubmit} fields={props.fields} isLoading={props.isLoading} />}
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
