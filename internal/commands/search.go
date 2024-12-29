package commands

import "github.com/bwmarrin/discordgo"

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

func SearchCommandHandler(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Okay here's art that I found matching your query",
		},
	})
}