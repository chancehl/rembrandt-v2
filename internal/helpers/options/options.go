package options

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func GetOption(i *discordgo.InteractionCreate, name string) (*discordgo.ApplicationCommandInteractionDataOption, error) {
	options := createOptionMap(i)
	if option, ok := options[name]; ok {
		return option, nil
	}
	return nil, fmt.Errorf("option %s not found", name)
}

func createOptionMap(i *discordgo.InteractionCreate) map[string]*discordgo.ApplicationCommandInteractionDataOption {
	options := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(i.ApplicationCommandData().Options))
	for _, option := range i.ApplicationCommandData().Options {
		options[option.Name] = option
	}
	return options
}
