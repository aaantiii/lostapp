package models

import "time"

type Notification struct {
	ID              uint      `json:"id,omitempty" gorm:"primaryKey;autoIncrement"`
	CreatedAt       time.Time `json:"createdAt"`
	Title           string    `json:"title" gorm:"not null;size:128"`
	Message         string    `json:"message" gorm:"not null;size:2048"`
	Link            string    `json:"link,omitempty" gorm:"size:256"`
	SenderName      string    `json:"senderName" gorm:"not null;size:64"`
	SenderDiscordID string    `json:"senderDiscordID,omitempty" gorm:"not null;size:18"`
}

func NewNotification(title, message, link, senderName, senderDiscordID string) *Notification {
	return &Notification{
		Title:           title,
		Message:         message,
		Link:            link,
		SenderName:      senderName,
		SenderDiscordID: senderDiscordID,
	}
}
