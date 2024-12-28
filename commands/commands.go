package commands

import "github.com/bwmarrin/discordgo"

var CommandDefinitions = []*discordgo.ApplicationCommand{
	&ArtCommandDefinition,
	&SearchArtCommandDefinition,
	&SubscribeCommandDefinition,
}
