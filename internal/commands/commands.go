package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/chancehl/rembrandt-v2/internal/models"
)

var Commands = []*discordgo.ApplicationCommand{
	&ArtCommand,
	&SearchCommand,
	&SubscribeCommand,
}

var Handlers = map[string]func(*models.BotSession, *models.BotInteraction){
	"art":        ArtCommandHandler,
	"search-art": SearchCommandHandler,
	"subscribe":  SubscribeCommandHandler,
}
