package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/chancehl/rembrandt-v2/internal/context"
	"github.com/chancehl/rembrandt-v2/internal/utils"
)

// `/search` command definition
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

// `/search` command handler
func SearchCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate, ctx *context.BotContext) {
	query, _ := utils.GetOption(i.Interaction, "query")

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("You queried for: %s", query.StringValue()),
		},
	})
}
