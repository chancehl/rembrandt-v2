package commands

import (
	"github.com/bwmarrin/discordgo"
)

var ArtCommand = discordgo.ApplicationCommand{
	Name:        "art",
	Description: "Get a random piece of art",
}

func ArtCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
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
