package util

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

func MentionUserID(id string) string {
	return "<@!" + id + ">"
}

func MentionRole(roleID string) string {
	return "<@&" + roleID + ">"
}

func DeleteInteractionResponseWithTimeout(s *discordgo.Session, i *discordgo.Interaction, timeout time.Duration) error {
	time.Sleep(timeout)
	return s.InteractionResponseDelete(i)
}

func CreateMessageURL(guildID, channelID, messageID string) string {
	return fmt.Sprintf("https://discord.com/channels/%s/%s/%s\n", guildID, channelID, messageID)
}
