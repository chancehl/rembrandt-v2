package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/chancehl/rembrandt-v2/internal/clients/api/met"
	"github.com/chancehl/rembrandt-v2/internal/utils"
)

// `/art` command definition
var ArtCommand = discordgo.ApplicationCommand{
	Name:        "art",
	Description: "Get a random piece of art",
}

// `/art` command handler
func ArtCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate, c *met.Client) {
	objectData, err := c.GetRandomObject()
	if err != nil {
		utils.RespondWithError(s, i, err)
		return
	}

	fields := []*discordgo.MessageEmbedField{}

	if objectData.ArtistDisplayName != "" {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "Artist",
			Value:  objectData.ArtistDisplayName,
			Inline: false,
		})
	}

	if objectData.ObjectDate != "" {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "Date",
			Value:  objectData.ObjectDate,
			Inline: false,
		})
	}

	if objectData.Department != "" {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "Department",
			Value:  objectData.Department,
			Inline: false,
		})
	}

	if objectData.Medium != "" {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "Medium",
			Value:  objectData.Medium,
			Inline: false,
		})
	}

	if objectData.AccessionNumber != "" {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "Accession Number",
			Value:  objectData.AccessionNumber,
			Inline: false,
		})
	}

	if objectData.ObjectURL != "" {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "URL",
			Value:  fmt.Sprintf("[View on The Met website](%s)", objectData.ObjectURL),
			Inline: false,
		})
	}

	embed := discordgo.MessageEmbed{
		Title:  objectData.Title,
		Fields: fields,
		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("The Metropolitan Museum of Art (Object ID: %d)", objectData.ObjectID),
		},
	}

	if objectData.PrimaryImage != "" {
		embed.Image = &discordgo.MessageEmbedImage{
			URL: objectData.PrimaryImage,
		}
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				&embed,
			},
		},
	})
}
