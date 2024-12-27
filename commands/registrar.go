package commands

import "github.com/bwmarrin/discordgo"

type CommandRegistrar interface {
	RegisterCommands([]*discordgo.ApplicationCommand) error
	DeregisterCommands([]*discordgo.ApplicationCommand) error
}
