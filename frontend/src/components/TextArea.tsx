import '@styles/components/TextArea.scss'
import { InputHTMLAttributes, forwardRef } from 'react'

type TextAreaProps = InputHTMLAttributes<HTMLTextAreaElement>

const TextArea = forwardRef<HTMLTextAreaElement, TextAreaProps>((props, forwardedRef) => (
  <textarea ref={forwardedRef} {...props} className={`TextArea ${props.className}`} />
))

export default TextArea
