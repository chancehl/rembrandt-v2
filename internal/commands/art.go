package commands

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/chancehl/rembrandt-v2/internal/context"
)

var ErrorMessage string = "Sorry, something went wrong when I was looking up your art. Please try again later."

// `/art` command definition
var ArtCommand = discordgo.ApplicationCommand{
	Name:        "art",
	Description: "Get a random piece of art",
}

// `/art` command handler
func ArtCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate, ctx *context.BotContext) {
	// tell the user the bot is "thinking..." while we wait for a completion
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})

	objectData, err := ctx.Clients.Met.GetRandomObject()

	if err != nil {
		log.Printf("error getting random object: %v", err)

		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Content: &ErrorMessage,
		})

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

	if summary, err := ctx.Clients.OpenAI.CreateSummaryForObject(objectData); err == nil {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "Description",
			Value:  summary.Description,
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

	if objectData.Culture != "" {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "Culture",
			Value:  objectData.Culture,
			Inline: false,
		})
	}

	if objectData.Period != "" {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "Period",
			Value:  objectData.Period,
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

	// edit the original message with the response
	s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Embeds: &[]*discordgo.MessageEmbed{
			&embed,
		},
	})
}
