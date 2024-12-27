package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/chancehl/rembrandt-v2/config"
)

type CommandRegistrar interface {
	RegisterCommands([]*discordgo.ApplicationCommand) error
	DeregisterCommands([]*discordgo.ApplicationCommand) error
}

type SlashCommandRegistrar struct {
	config     config.BotConfig
	session    *discordgo.Session
	commands   []*discordgo.ApplicationCommand
	registered []*discordgo.ApplicationCommand
}

func NewSlashCommandRegistrar(config config.BotConfig, session *discordgo.Session, commands []*discordgo.ApplicationCommand) *SlashCommandRegistrar {
	return &SlashCommandRegistrar{
		config:     config,
		session:    session,
		commands:   commands,
		registered: []*discordgo.ApplicationCommand{},
	}
}

func (r *SlashCommandRegistrar) RegisterCommands() error {
	for _, cmd := range r.commands {
		if _, err := r.session.ApplicationCommandCreate(r.session.State.User.ID, r.config.TestGuildID, cmd); err != nil {
			return fmt.Errorf("cannot register command %s: %v", cmd.Name, err)
		}
	}
	return nil
}

func (r *SlashCommandRegistrar) DeregisterCommands() error {
	if r.config.RemoveCommandsOnExit {
		for _, cmd := range r.registered {
			err := r.session.ApplicationCommandDelete(r.session.State.User.ID, r.config.TestGuildID, cmd.ID)
			if err != nil {
				return fmt.Errorf("cannot deregister command %s: %v", cmd.Name, err)
			}
		}
	}
	return nil
}
