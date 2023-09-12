import '@styles/components/Dialog.scss'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import Button from './Button'
import * as d from '@radix-ui/react-dialog'
import { faXmark } from '@fortawesome/free-solid-svg-icons'
import Form, { FormProps } from './Form'
import { useState } from 'react'

interface DialogProps extends FormProps {
  title: string
  description: string
  onOpenChange?: (open: boolean) => void
  onClose?: () => void
}

export default function Dialog(props: DialogProps) {
  const [openState, setOpenState] = useState(false)

  function handleSubmit(data: any) {
    props.onSubmit(data)
    setOpenState(false)
  }

  return (
    <d.Root open={openState} onOpenChange={props.onOpenChange}>
      <d.Trigger asChild>
        <Button onClick={() => setOpenState(true)}>{props.title}</Button>
      </d.Trigger>
      <d.Portal>
        <d.Overlay className="DialogOverlay" />
        <d.Content className="DialogContent">
          <d.Title className="title">{props.title}</d.Title>
          <d.Description className="description">{props.description}</d.Description>
          <Form submitText={props.submitText} onSubmit={handleSubmit} fields={props.fields}></Form>
          <d.Close asChild className="close">
            <button aria-label="Schließen" onClick={() => setOpenState(false)}>
              <FontAwesomeIcon icon={faXmark} />
            </button>
          </d.Close>
        </d.Content>
      </d.Portal>
    </d.Root>
  )
}
