package types

type ApiMessage struct {
	Message string `json:"message"`
}

func NewMessage(msg string) ApiMessage {
	return ApiMessage{
		Message: msg,
	}
}
