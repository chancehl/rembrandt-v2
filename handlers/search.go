package handlers

import "github.com/bwmarrin/discordgo"

func SearchArtCommandHandler(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Okay here's art that I found matching your query",
		},
	})
}
