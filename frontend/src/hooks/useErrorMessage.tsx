import { HttpStatusCode } from 'axios'

type ErrorCode = number | undefined

const httpErrorMessages: ReadonlyMap<ErrorCode, string> = new Map([
  [HttpStatusCode.BadRequest, 'Der Server hat eine ungültige Anfrage erhalten.'],
  [HttpStatusCode.Unauthorized, 'Du bist nicht angemeldet.'],
  [HttpStatusCode.Forbidden, 'Du hast keine Berechtigung für diese Aktion.'],
  [HttpStatusCode.NotFound, 'Die Angeforderte Ressource wurde nicht gefunden.'],
  [HttpStatusCode.RequestTimeout, 'Die Antwort des Servers hat zu lange gedauert.'],
  [HttpStatusCode.Conflict, 'Bei der Anfrage ist ein unerwarteter Konflikt aufgetreten.'],
  [HttpStatusCode.InternalServerError, 'Beim bearbeiten der Anfrage ist ein Fehler aufgetreten.'],
])

export default function useErrorMessage(code: ErrorCode) {
  const message = httpErrorMessages.get(code) ?? 'Ein unbekannter Fehler ist aufgetreten.'

  return <p className="error">{message}</p>
}
