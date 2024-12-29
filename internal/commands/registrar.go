package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/chancehl/rembrandt-v2/config"
	"github.com/chancehl/rembrandt-v2/internal/models"
)

// type commandRegistrar interface {
// 	RegisterCommands([]*discordgo.ApplicationCommand) error
// 	DeregisterCommands([]*discordgo.ApplicationCommand) error
// }

type CommandRegistrar struct {
	config     config.BotConfig
	session    *discordgo.Session
	commands   []*discordgo.ApplicationCommand
	handlers   map[string]func(*models.BotSession, *models.BotInteraction)
	registered []*discordgo.ApplicationCommand
}

func NewCommandRegistrar(config config.BotConfig, session *discordgo.Session) *CommandRegistrar {
	return &CommandRegistrar{
		config:     config,
		session:    session,
		commands:   Commands,
		handlers:   Handlers,
		registered: []*discordgo.ApplicationCommand{},
	}
}

// Registers commands on bot startup
func (r *CommandRegistrar) RegisterCommands() error {
	for i := range r.commands {
		cmd, err := r.session.ApplicationCommandCreate(r.session.State.User.ID, r.config.TestGuildID, r.commands[i])
		if err != nil {
			return fmt.Errorf("cannot register command %s: %v", cmd.Name, err)
		}
		r.session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if handler, ok := Handlers[i.ApplicationCommandData().Name]; ok {
				session := models.NewBotSession(s)
				interaction := models.NewBotInteraction(i)
				handler(session, interaction)
			}
		})
		r.registered = append(r.registered, cmd)
	}
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
