import '@styles/components/Form.scss'
import * as f from '@radix-ui/react-form'
import Button from './Button'

export interface FormProps {
  fields: (FormField | null)[]
  submitText: string
  isLoading?: boolean
  onSubmit: (data: any) => void
}

export interface FormField {
  label: string
  messages?: Pick<f.FormMessageProps, 'match' | 'children' | 'id'>[]
  name: string
  control: JSX.Element
  type?: 'text' | 'number' | 'date'
  noSubmit?: boolean
}

export default function Form({ fields, submitText, isLoading, onSubmit }: FormProps) {
  function handleSubmit(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault()

    const formData = new FormData(e.currentTarget)
    let data: any = {}
    for (const field of fields) {
      if (field === null || field.noSubmit) continue
      const formValue = formData.get(field.name)?.toString() ?? ''
      switch (field.type) {
        case 'number':
          data[field.name] = formValue !== '' ? parseInt(formValue) : ''
          break
        case 'date':
          data[field.name] = formValue !== '' ? new Date(formValue).toISOString() : ''
          break
        default:
          data[field.name] = formValue.trim().replaceAll(/\s\s+/g, ' ')
      }

      if (data[field.name] === '') return
    }

    onSubmit(data)
  }

  return (
    <div className="Form">
      <f.Root className="root" onSubmit={handleSubmit}>
        {fields.map(
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
        )}
        <f.Submit asChild>
          <Button className="submit-button" isLoading={isLoading}>
            {submitText}
          </Button>
        </f.Submit>
      </f.Root>
    </div>
  )
}
