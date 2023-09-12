import '@styles/components/Form.scss'
import * as f from '@radix-ui/react-form'
import Button from './Button'

export interface FormProps {
  fields: FormField[]
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
}

export default function Form({ fields, submitText, isLoading, onSubmit }: FormProps) {
  function handleSubmit(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault()

    const formData = new FormData(e.currentTarget)
    let data: any = {}
    for (const field of fields) {
      const formValue = formData.get(field.name)?.toString() ?? ''
      switch (field.type) {
        case 'number':
          data[field.name] = parseInt(formValue)
          break
        case 'date':
          data[field.name] = new Date(formValue)
          break
        default:
          data[field.name] = formValue.trim().replaceAll(/\s\s+/g, ' ')
      }
    }

    onSubmit(data)
  }

  return (
    <div className="Form">
      <f.Root className="Root" onSubmit={handleSubmit}>
        {fields.map((field) => (
          <f.Field key={field.name} className="Field" name={field.name}>
            <f.Label className="Label">{field.label}:</f.Label>
            <f.Control asChild className="Control">
              {field.control}
            </f.Control>
            {field.messages &&
              field.messages.map((message) => (
                <f.Message key={message.id} className="Message" match={message.match}>
                  {message.children}
                </f.Message>
              ))}
          </f.Field>
        ))}
        <f.Submit asChild>
          <Button className="SubmitButton" disabled={isLoading}>
            {submitText}
          </Button>
        </f.Submit>
      </f.Root>
    </div>
  )
}
