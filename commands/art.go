package commands

import "github.com/bwmarrin/discordgo"

var ArtCommandDefinition = discordgo.ApplicationCommand{
	Name:        "art",
	Description: "Get a random piece of art",
}
