import '@styles/components/UnsubmittableForm.scss'

interface UnsubmittableFormProps {
  children: JSX.Element[] | JSX.Element
}

export default function UnsubmittableForm({ children }: UnsubmittableFormProps) {
  return (
    <div className="UnsubmittableForm">
      <form onSubmit={(e) => e.preventDefault()}>{children}</form>
    </div>
  )
}
