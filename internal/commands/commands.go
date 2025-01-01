package commands

import (
	"github.com/bwmarrin/discordgo"
)

var Commands = []*discordgo.ApplicationCommand{
	&ArtCommand,
	&SearchCommand,
	&SubscribeCommand,
	&UnsubscribeCommand,
}

var Handlers = map[string]HandlerFunc{
	ArtCommand.Name:         ArtCommandHandler,
	SearchCommand.Name:      SearchCommandHandler,
	SubscribeCommand.Name:   SubscribeCommandHandler,
	UnsubscribeCommand.Name: UnsubscribeCommandHandler,
}
