package models

import "github.com/bwmarrin/discordgo"

type User struct {
	DiscordID string `gorm:"size:18;primaryKey;not null"`
	Name      string `gorm:"size:32;not null"`
	IsAdmin   bool   `gorm:"default:false;not null"`
}

func (u *User) Mention() string {
	return "<@" + u.DiscordID + ">"
}

func UserFromGuildMember(member *discordgo.Member) *User {
	return &User{
		DiscordID: member.User.ID,
		Name:      member.Nick,
	}
}
