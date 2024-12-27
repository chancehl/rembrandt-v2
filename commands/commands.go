package commands

import "github.com/bwmarrin/discordgo"

var AllCommands = []*discordgo.ApplicationCommand{
	&ArtCommand,
	&SearchArtCommand,
	&SubscribeCommand,
}
