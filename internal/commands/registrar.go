package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/chancehl/rembrandt-v2/internal/context"
)

// type commandRegistrar interface {
// 	RegisterCommands([]*discordgo.ApplicationCommand) error
// 	DeregisterCommands([]*discordgo.ApplicationCommand) error
// }

type HandlerFunc func(*discordgo.Session, *discordgo.InteractionCreate, *context.AppContext)

type Registrar struct {
	commands   []*discordgo.ApplicationCommand
	registered []*discordgo.ApplicationCommand
	handlers   map[string]HandlerFunc
	context    *context.AppContext
}

func NewRegistrar(context *context.AppContext) *Registrar {
	return &Registrar{
		commands:   Commands,
		handlers:   Handlers,
		registered: []*discordgo.ApplicationCommand{},
		context:    context,
	}
}

// Registers commands on bot startup
func (r *Registrar) RegisterCommands() error {
	for idx := range r.commands {
		userID := r.context.Session.State.User.ID
		testGuildID := r.context.Config.TestGuildID

		cmd, err := r.context.Session.ApplicationCommandCreate(userID, testGuildID, r.commands[idx])
		if err != nil {
			return fmt.Errorf("cannot register command %s: %v", cmd.Name, err)
		}

		r.registered = append(r.registered, cmd)
	}

	r.context.Session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if handler, ok := Handlers[i.ApplicationCommandData().Name]; ok {
			handler(s, i, r.context)
		}
	})

	return nil
}

// De-registers commands on bot exit
func (r *Registrar) DeregisterCommands() error {
	if r.context.Config.RemoveCommandsOnExit {
		for _, cmd := range r.registered {
			userID := r.context.Session.State.User.ID
			testGuildID := r.context.Config.TestGuildID

			err := r.context.Session.ApplicationCommandDelete(userID, testGuildID, cmd.ID)
			if err != nil {
				return fmt.Errorf("cannot deregister command %s: %v", cmd.Name, err)
			}
		}
	}
	return nil
}
