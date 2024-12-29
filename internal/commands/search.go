package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/chancehl/rembrandt-v2/internal/models"
)

var SearchCommand = discordgo.ApplicationCommand{
	Name:        "search-art",
	Description: "Searches the MET for a piece of art",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "query",
			Description: "The search query",
			Type:        discordgo.ApplicationCommandOptionString,
			Required:    true,
		},
	},
}

func SearchCommandHandler(session *models.BotSession, interaction *models.BotInteraction) {
	session.RespondToInteractionWithString(interaction, "Okay, here's the art you searched for")
}
