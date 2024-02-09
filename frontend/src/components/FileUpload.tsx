import '@styles/components/FileUpload.scss'
import { forwardRef, useId, useRef, useState } from 'react'
import Button from './Button'

type FileUploadProps = Omit<React.InputHTMLAttributes<HTMLInputElement>, 'title'> & {
  title: string
}

const FileUpload = forwardRef<HTMLInputElement, FileUploadProps>(({ title, id, ...inputProps }, forwardedRef) => {
  const newId = useId()
  if (!id) id = newId

  const inputRef = useRef<HTMLInputElement | null>(null)
  const [fileName, setFileName] = useState('')

  function handleClick() {
    inputRef.current?.click()
  }

  function handleChange(e: React.ChangeEvent<HTMLInputElement>) {
    if (e.currentTarget.files?.[0] !== undefined) {
      setFileName(e.currentTarget.files[0].name)
    }
  }

  function refCallback(node: HTMLInputElement) {
    inputRef.current = node

    if (typeof forwardedRef === 'function') {
      forwardedRef(node)
    } else if (forwardedRef) {
      forwardedRef.current = node
    }
  }

  return (
    <div className="FileUpload">
      <div className="wrapper">
        <Button onClick={handleClick}>{title}</Button>
        {fileName && <span>{fileName}</span>}
      </div>
      <input {...inputProps} type="file" ref={refCallback} onChange={handleChange} />
    </div>
  )
})

export default FileUpload
