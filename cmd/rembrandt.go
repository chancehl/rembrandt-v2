package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/chancehl/rembrandt-v2/internal/cache"
	"github.com/chancehl/rembrandt-v2/internal/clients/db"
	"github.com/chancehl/rembrandt-v2/internal/clients/met"
	"github.com/chancehl/rembrandt-v2/internal/clients/openai"
	"github.com/chancehl/rembrandt-v2/internal/commands"
	"github.com/chancehl/rembrandt-v2/internal/config"
	internal "github.com/chancehl/rembrandt-v2/internal/context"
	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

var (
	cfg       config.Config
	ctx       *internal.BotContext
	registrar *commands.Registrar
)

func init() {
	// load dot env variables
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	// create config
	if err := envconfig.Process(context.TODO(), &cfg); err != nil {
		log.Fatalf("could not load context from environment variables: %v", err)
	}

	// create bot session
	session, err := discordgo.New("Bot " + cfg.Discord.BotToken)
	if err != nil {
		log.Fatalf("invalid bot parameters: %v", err)
	}

	// create cache
	inMemoryCache := cache.NewInMemoryCache() // TODO: convert to singleton

	// create clients
	metClient := met.NewClient(inMemoryCache)
	openAIClient := openai.NewClient(cfg.OpenAI.Key, inMemoryCache)
	dbClient, err := db.NewClient(cfg.DB.URL, inMemoryCache)
	if err != nil {
		log.Fatalf("could not create db client: %v", err)
	}

	clients := &internal.Clients{
		Met:    metClient,
		DB:     dbClient,
		OpenAI: openAIClient,
	}

	// create context
	ctx = &internal.BotContext{
		Config:  &cfg,
		Clients: clients,
		Session: session,
	}

	// create command registrar
	registrar = commands.NewRegistrar(ctx)
}

func main() {
	log.Println("starting bot...")

	// start the bot
	if err := ctx.Session.Open(); err != nil {
		log.Fatalf("cannot open the session: %v", err)
	}
	defer ctx.Session.Close()

	// hydrate cache on startup
	log.Println("hydrating cache...")
	if resp, err := ctx.Clients.Met.GetObjectIDs(); err == nil {
		log.Printf("successfully fetched %d object IDs from met api", resp.Total)
	} else {
		log.Fatalf("failed to hydrate cache with initial data: %v", err)
	}

	// register commands
	log.Printf("registering %d bot command(s)...", len(commands.Commands))
	if err := registrar.RegisterCommands(); err != nil {
		log.Fatalf("cannot register commands: %v", err)
	} else {
		log.Println("successfully registered commands")
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("bot has started (press ctrl+c to exit)")
	<-stop

	// cleanup
	if err := registrar.DeregisterCommands(); err != nil {
		log.Fatalf("cannot deregister commands: %v", err)
	} else {
		log.Println("successfully deregistered commands")
	}

	log.Println("bot exited gracefully")
}
