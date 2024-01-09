package util

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

func MentionRole(roleID string) string {
	return "<@&" + roleID + ">"
}

func DeleteInteractionResponseWithTimeout(s *discordgo.Session, i *discordgo.Interaction, timeout time.Duration) error {
	time.Sleep(timeout)
	return s.InteractionResponseDelete(i)
}
