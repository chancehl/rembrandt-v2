package handlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/chancehl/rembrandt-v2/commands"
)

var CommandHandlers = map[string]func(*discordgo.Session, *discordgo.InteractionCreate){
	commands.ArtCommandDefinition.Name:       ArtCommandHandler,
	commands.SearchArtCommandDefinition.Name: SearchArtCommandHandler,
	commands.SubscribeCommandDefinition.Name: SubscribeCommandHandler,
}
