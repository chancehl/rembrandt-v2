package commands

import (
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
	}

	summary, err := ctx.Clients.OpenAI.CreateSummaryForObject(objectData)
	if err != nil {
		log.Printf("could not generate embed for object ID %d: %v", objectData.ObjectID, err)
		objectData.Summary = objectData.ObjectName // fallback to object name if we can't generate a summary
	} else {
		objectData.Summary = summary.Description
	}

	// generate embed
	embed := objectData.GenerateEmbed()

	// edit the original message with the response
	s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Embeds: &[]*discordgo.MessageEmbed{
			embed,
		},
	})
}
