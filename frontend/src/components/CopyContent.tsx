import '@styles/components/CopyContent.scss'
import { faCopy } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { useMessages } from '@context/messagesContext'
import { useTranslation } from 'react-i18next'

type CopyContentProps = {
  children: string
}

export default function CopyContent({ children }: CopyContentProps) {
  const [t] = useTranslation('components/copyContent')
  const { sendMessage } = useMessages()

  function handleCopy() {
    navigator.clipboard.writeText(children)
    sendMessage({ type: 'success', message: t('success') })
  }

  return (
    <div className="CopyContent" role="button" onClick={handleCopy} title="Click to copy" aria-description="Click to copy link">
      <pre role="link">{children}</pre>
      <FontAwesomeIcon icon={faCopy} className="icon" />
    </div>
  )
}
