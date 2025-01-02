package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/chancehl/rembrandt-v2/internal/clients/api/met"
	"github.com/chancehl/rembrandt-v2/internal/config"
)

// type commandRegistrar interface {
// 	RegisterCommands([]*discordgo.ApplicationCommand) error
// 	DeregisterCommands([]*discordgo.ApplicationCommand) error
// }

type HandlerFunc func(s *discordgo.Session, i *discordgo.InteractionCreate, mc *met.METAPIClient)

type CommandRegistrar struct {
	config     config.BotConfig
	session    *discordgo.Session
	commands   []*discordgo.ApplicationCommand
	registered []*discordgo.ApplicationCommand
	handlers   map[string]HandlerFunc
	client     *met.METAPIClient
}

func NewCommandRegistrar(config config.BotConfig, session *discordgo.Session, client *met.METAPIClient) *CommandRegistrar {
	return &CommandRegistrar{
		config:     config,
		session:    session,
		commands:   Commands,
		handlers:   Handlers,
		registered: []*discordgo.ApplicationCommand{},
		client:     client,
	}
}

// Registers commands on bot startup
func (r *CommandRegistrar) RegisterCommands() error {
	for idx := range r.commands {
		cmd, err := r.session.ApplicationCommandCreate(r.session.State.User.ID, r.config.TestGuildID, r.commands[idx])
		if err != nil {
			return fmt.Errorf("cannot register command %s: %v", cmd.Name, err)
		}
		r.registered = append(r.registered, cmd)
	}
	r.session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if handler, ok := Handlers[i.ApplicationCommandData().Name]; ok {
			handler(s, i, r.client)
		}
	})
	return nil
}

// De-registers commands on bot exit
func (r *CommandRegistrar) DeregisterCommands() error {
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
