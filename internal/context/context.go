package context

import (
	"github.com/chancehl/rembrandt-v2/internal/clients/db"
	"github.com/chancehl/rembrandt-v2/internal/clients/met"
	"github.com/chancehl/rembrandt-v2/internal/clients/openai"
)

type AppContext struct {
	MetClient    *met.Client
	DbClient     *db.Client
	OpenAIClient *openai.Client
}
