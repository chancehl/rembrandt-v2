package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/chancehl/rembrandt-v2/internal/helpers/interaction"
	"github.com/chancehl/rembrandt-v2/internal/helpers/responses"
)

var SearchCommand = discordgo.ApplicationCommand{
	Name:        "search",
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

func SearchCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	query, _ := interaction.GetOption(i.Interaction, "query")
	responses.RespondWithString(s, i.Interaction, fmt.Sprintf("You queried for: %s", query.StringValue()))
}
