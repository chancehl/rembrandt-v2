package context

import (
	"github.com/bwmarrin/discordgo"
	"github.com/chancehl/rembrandt-v2/internal/clients/db"
	"github.com/chancehl/rembrandt-v2/internal/clients/met"
	"github.com/chancehl/rembrandt-v2/internal/clients/openai"
	"github.com/chancehl/rembrandt-v2/internal/config"
)

type ClientContext struct {
	Met    *met.Client
	DB     *db.Client
	OpenAI *openai.Client
}

func NewClientContext(met *met.Client, db *db.Client, openai *openai.Client) *ClientContext {
	return &ClientContext{Met: met, DB: db, OpenAI: openai}
}

type AppContext struct {
	Clients *ClientContext
	Config  *config.Config
	Session *discordgo.Session
}

func NewAppContext(clients *ClientContext, config *config.Config, session *discordgo.Session) *AppContext {
	return &AppContext{Clients: clients, Config: config, Session: session}
}
