import '@styles/components/TextArea.scss'
import { forwardRef } from 'react'

type TextAreaProps = React.InputHTMLAttributes<HTMLTextAreaElement>

const TextArea = forwardRef((props: TextAreaProps, ref: React.Ref<HTMLTextAreaElement>) => (
  <textarea ref={ref} {...props} className={`TextArea ${props.className}`} />
))

export default TextArea
