package commands

import "github.com/bwmarrin/discordgo"

var SearchArtCommandDefinition = discordgo.ApplicationCommand{
	Name:        "search-art",
	Description: "Searches the MET for a piece of art",
}
