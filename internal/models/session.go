package models

import "github.com/bwmarrin/discordgo"

type BotSession struct {
	*discordgo.Session
}

func NewBotSession(s *discordgo.Session) *BotSession {
	return &BotSession{s}
}

func (s *BotSession) RespondToInteractionWithString(i *BotInteraction, content string) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
		},
	})
}

func (s *BotSession) RespondToInteractionWithEmbed(i *BotInteraction, embeds []*discordgo.MessageEmbed) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: embeds,
		},
	})
}
