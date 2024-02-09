import '@styles/components/Form.scss'
import * as f from '@radix-ui/react-form'
import Button from './Button'
import { FormEvent, createRef } from 'react'

export type FormProps = {
  children: (FormField | null)[] | FormField
  submitText: string
  loading?: boolean
  onSubmit?: (data: FormData) => boolean | void | Promise<boolean | void>
  submitControl?: JSX.Element
}

export type FormField = {
  label: string
  messages?: Pick<f.FormMessageProps, 'match' | 'children' | 'id'>[]
  name: string
  control: JSX.Element
}

export default function Form({ children, submitText, loading, onSubmit, submitControl }: FormProps) {
  function handleSubmit(e: FormEvent<HTMLFormElement>) {
    if (submitControl) {
      e.stopPropagation()
      return
    }

    e.preventDefault()
    const data = new FormData(e.currentTarget)
    onSubmit?.(data)
  }

  return (
    <div className="Form">
      <f.Root className="root" onSubmit={handleSubmit}>
        {Array.isArray(children) ? (
          children.map(
            (field) =>
              field && (
                <f.Field key={field.name} className="field" name={field.name}>
                  <f.Label className="label">{field.label}</f.Label>
                  <f.Control asChild className="control">
                    {field.control}
                  </f.Control>
                  {field.messages &&
                    field.messages.map((message) => (
                      <f.Message key={message.id} className="message" match={message.match}>
                        {message.children}
                      </f.Message>
                    ))}
                </f.Field>
              )
          )
        ) : (
          <f.Field className="field" name={children.name}>
            <f.Label className="label">{children.label}</f.Label>
            <f.Control asChild className="control">
              {children.control}
            </f.Control>
            {children.messages &&
              children.messages.map((message) => (
                <f.Message key={message.id} className="message" match={message.match}>
                  {message.children}
                </f.Message>
              ))}
          </f.Field>
        )}
        <f.Submit asChild>
          {submitControl ?? (
            <Button className="submit-button" loading={loading}>
              {submitText}
            </Button>
          )}
        </f.Submit>
      </f.Root>
    </div>
  )
}
