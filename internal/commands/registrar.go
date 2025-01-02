package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/chancehl/rembrandt-v2/internal/config"
	"github.com/chancehl/rembrandt-v2/internal/context"
)

// type commandRegistrar interface {
// 	RegisterCommands([]*discordgo.ApplicationCommand) error
// 	DeregisterCommands([]*discordgo.ApplicationCommand) error
// }

type HandlerFunc func(*discordgo.Session, *discordgo.InteractionCreate, *context.AppContext)

type Registrar struct {
	config     *config.Config
	session    *discordgo.Session
	commands   []*discordgo.ApplicationCommand
	registered []*discordgo.ApplicationCommand
	handlers   map[string]HandlerFunc
	appContext *context.AppContext
}

func NewRegistrar(config *config.Config, session *discordgo.Session, appContext *context.AppContext) *Registrar {
	return &Registrar{
		config:     config,
		session:    session,
		commands:   Commands,
		handlers:   Handlers,
		registered: []*discordgo.ApplicationCommand{},
		appContext: appContext,
	}
}

// Registers commands on bot startup
func (r *Registrar) RegisterCommands() error {
	for idx := range r.commands {
		cmd, err := r.session.ApplicationCommandCreate(r.session.State.User.ID, r.config.TestGuildID, r.commands[idx])
		if err != nil {
			return fmt.Errorf("cannot register command %s: %v", cmd.Name, err)
		}
		r.registered = append(r.registered, cmd)
	}
	r.session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if handler, ok := Handlers[i.ApplicationCommandData().Name]; ok {
			handler(s, i, r.appContext)
		}
	})
	return nil
}

// De-registers commands on bot exit
func (r *Registrar) DeregisterCommands() error {
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
