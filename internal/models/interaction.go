package models

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type BotInteraction struct {
	*discordgo.InteractionCreate
}

func NewBotInteraction(i *discordgo.InteractionCreate) *BotInteraction {
	return &BotInteraction{i}
}

func (i *BotInteraction) GetOption(name string) (*discordgo.ApplicationCommandInteractionDataOption, error) {
	options := createOptionMap(i.Interaction)
	if option, ok := options[name]; ok {
		return option, nil
	}
	return nil, fmt.Errorf("option %s not found", name)
}

func createOptionMap(i *discordgo.Interaction) map[string]*discordgo.ApplicationCommandInteractionDataOption {
	options := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(i.ApplicationCommandData().Options))
	for _, option := range i.ApplicationCommandData().Options {
		options[option.Name] = option
	}
	return options
}
