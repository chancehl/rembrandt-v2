package context

import (
	"github.com/bwmarrin/discordgo"
	"github.com/chancehl/rembrandt-v2/internal/cache"
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

type AppContext struct {
	Clients *ClientContext
	Config  *config.Config
	Session *discordgo.Session
	Cache   *cache.InMemoryCache
}
