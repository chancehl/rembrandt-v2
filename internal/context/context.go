package context

import (
	"github.com/chancehl/rembrandt-v2/internal/clients/db"
	"github.com/chancehl/rembrandt-v2/internal/clients/met"
	"github.com/chancehl/rembrandt-v2/internal/clients/openai"
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
}

func NewAppContext(clients *ClientContext) *AppContext {
	return &AppContext{Clients: clients}
}
