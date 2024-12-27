package commands

import "github.com/bwmarrin/discordgo"

var ArtCommand = discordgo.ApplicationCommand{
	Name:        "art",
	Description: "Get a random piece of art",
}
