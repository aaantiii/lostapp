import '../scss/components/Dialog.scss'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import Button from './Button'
import * as d from '@radix-ui/react-dialog'
import { faXmark } from '@fortawesome/free-solid-svg-icons'

interface DialogProps {
  title: string
  description?: string
  fieldsets?: JSX.Element | JSX.Element[]
  open?: boolean
  onOpenChange?: (open: boolean) => void
}

export default function Dialog({ title, description, fieldsets, open, onOpenChange }: DialogProps) {
  return (
    <div className="Dialog">
      <d.Root open={open} onOpenChange={onOpenChange}>
        <d.Trigger asChild>
          <Button>Schließen</Button>
        </d.Trigger>
        <d.Portal>
          <d.Overlay className="DialogOverlay" />
          <d.Content className="DialogContent">
            <d.Title className="DialogTitle">Edit profile</d.Title>
            <d.Description className="DialogDescription">Make changes to your profile here. Click save when you're done.</d.Description>
            {fieldsets && <div className="DialogFieldsets">{fieldsets}</div>}
            <div style={{ display: 'flex', marginTop: 25, justifyContent: 'flex-end' }}>
              <d.Close asChild>
                <button className="Button green">Save changes</button>
              </d.Close>
            </div>
            <d.Close asChild>
              <button className="IconButton" aria-label="Close">
                <FontAwesomeIcon icon={faXmark} />
              </button>
            </d.Close>
          </d.Content>
        </d.Portal>
      </d.Root>
    </div>
  )
}
