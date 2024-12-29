package responses

import "github.com/bwmarrin/discordgo"

func RespondWithString(s *discordgo.Session, i *discordgo.Interaction, content string) {
	s.InteractionRespond(i, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
		},
	})
}

func RespondWithEmbed(s *discordgo.Session, i *discordgo.Interaction, embeds []*discordgo.MessageEmbed) {
	s.InteractionRespond(i, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: embeds,
		},
	})
}
