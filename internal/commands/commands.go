package commands

import (
	"github.com/bwmarrin/discordgo"
)

var Commands = []*discordgo.ApplicationCommand{
	&ArtCommand,
	&SearchCommand,
	&SubscribeCommand,
}

var Handlers = map[string]func(*discordgo.Session, *discordgo.InteractionCreate){
	"art":        ArtCommandHandler,
	"search-art": SearchCommandHandler,
	"subscribe":  SubscribeCommandHandler,
}
