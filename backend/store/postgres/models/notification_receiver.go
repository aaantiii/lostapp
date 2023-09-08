package models

type NotificationReceiver struct {
	Seen      bool   `json:"notificationRead" gorm:"default:false;not null"`
	DiscordID string `json:"discordID" gorm:"primaryKey;not null;size:18"`

	NotificationID uint `json:"notificationId" gorm:"primaryKey;not null"`
	Notification   *Notification
}
