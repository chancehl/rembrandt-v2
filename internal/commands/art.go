package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/chancehl/rembrandt-v2/internal/models"
)

var ArtCommand = discordgo.ApplicationCommand{
	Name:        "art",
	Description: "Get a random piece of art",
}

func ArtCommandHandler(session *models.BotSession, interaction *models.BotInteraction) {
	session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       "Art Title",
					Description: "Art Description",
					Fields: []*discordgo.MessageEmbedField{
						{
							Name:   "Artist",
							Value:  "Foo",
							Inline: false,
						},
					},
				},
			},
		},
	})
}
