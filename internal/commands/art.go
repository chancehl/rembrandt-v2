package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/chancehl/rembrandt-v2/internal/api"
)

// `/art` command definition
var ArtCommand = discordgo.ApplicationCommand{
	Name:        "art",
	Description: "Get a random piece of art",
}

// `/art` command handler
func ArtCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate, c *api.METAPIClient) {
	objectIDData, _ := c.GetObjectIDs()
	fmt.Println(objectIDData.Total)

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
