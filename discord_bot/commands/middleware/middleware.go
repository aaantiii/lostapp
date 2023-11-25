package middleware

import "github.com/bwmarrin/discordgo"

type InteractionMiddleware func(s *discordgo.Session, i *discordgo.InteractionCreate) error
