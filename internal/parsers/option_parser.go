package parsers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func createOptionMap(i *discordgo.Interaction) map[string]*discordgo.ApplicationCommandInteractionDataOption {
	options := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(i.ApplicationCommandData().Options))
	for _, option := range i.ApplicationCommandData().Options {
		options[option.Name] = option
	}
	return options
}

func GetOption(i *discordgo.Interaction, name string) (*discordgo.ApplicationCommandInteractionDataOption, error) {
	options := createOptionMap(i)
	if option, ok := options[name]; ok {
		return option, nil
	}
	return nil, fmt.Errorf("option %s not found", name)
}
