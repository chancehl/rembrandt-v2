package handlers

import "github.com/bwmarrin/discordgo"

func SubscribeCommandHandler(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Okay I should subscribe here",
		},
	})
}
