package handlers

import "github.com/bwmarrin/discordgo"

func ArtCommandHandler(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Here's a random piece of art!",
		},
	})
}
