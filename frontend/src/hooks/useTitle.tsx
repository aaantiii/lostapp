import { Helmet } from 'react-helmet-async'

export default function useTitle(title: string) {
  return (
    <>
      <Helmet title={title} />
      <h1>{title}</h1>
    </>
  )
}
