package utils

import "github.com/bwmarrin/discordgo"

func CreateOptionMap(i *discordgo.InteractionCreate) map[string]*discordgo.ApplicationCommandInteractionDataOption {
	options := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(i.ApplicationCommandData().Options))
	for _, option := range i.ApplicationCommandData().Options {
		options[option.Name] = option
	}
	return options
}
