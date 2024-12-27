package handlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/chancehl/rembrandt-v2/commands"
)

var CommandHandlers = map[string]func(*discordgo.Session, *discordgo.InteractionCreate){
	commands.ArtCommand.Name:       ArtCommandHandler,
	commands.SearchArtCommand.Name: SearchArtCommandHandler,
	commands.SubscribeCommand.Name: SubscribeCommandHandler,
}
