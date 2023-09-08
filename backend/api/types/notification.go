package types

type NotificationForm struct {
	Title              string               `json:"title" binding:"required" form:"title"`
	Message            string               `json:"message" binding:"required" form:"message"`
	Link               string               `json:"link" form:"link"`
	RecieverType       NotificationFormType `json:"recieverType" binding:"required" form:"recieverType"`
	RecieverIdentifier string               `json:"recieverIdentifier" form:"recieverIdentifier"`
}

type NotificationFormType string

const (
	NotificationFormTypeEveryone NotificationFormType = "everyone"
	NotificationFormTypeClan     NotificationFormType = "clan"
)
