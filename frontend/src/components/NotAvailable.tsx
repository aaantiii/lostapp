import AlertDialog from './AlertDialog'
import Button from './Button'

type NotAvailableProps = {
  trigger: React.ReactNode
}

export default function NotAvailable({ trigger }: NotAvailableProps) {
  return <AlertDialog title="Bald verfügbar" description="Diese Funktion ist noch nicht verfügbar." trigger={trigger} confirm={<Button>OK</Button>} />
}
