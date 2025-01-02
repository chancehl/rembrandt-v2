package main

import (
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
	"github.com/chancehl/rembrandt-v2/internal/context"
	"github.com/joho/godotenv"
)

var (
	ctx           *context.AppContext
	registrar     *commands.Registrar
	inMemoryCache *cache.InMemoryCache
)

func init() {
	// load dot env variables
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	// create bot session
	session, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		log.Fatalf("invalid bot parameters: %v", err)
	}

	// create cache
	inMemoryCache = cache.NewInMemoryCache()

	// create clients
	clients := &context.ClientContext{
		Met:    met.NewClient(inMemoryCache),
		DB:     db.NewClient(),
		OpenAI: openai.NewClient(),
	}

	// create config
	config := &config.Config{
		TestGuildID:          os.Getenv("TEST_GUILD_ID"),
		RemoveCommandsOnExit: os.Getenv("REMOVE_COMMANDS_ON_EXIT") == "true",
		HydrateCacheOnStart:  os.Getenv("HYDRATE_CACHE_ON_START") == "true",
	}

	// create context
	ctx = &context.AppContext{
		Clients: clients,
		Config:  config,
		Session: session,
	}

	// create command registrar
	registrar = commands.NewRegistrar(ctx)
}

func main() {
	log.Printf("starting bot with config %+v\n", *ctx.Config)

	// start the bot
	if err := ctx.Session.Open(); err != nil {
		log.Fatalf("cannot open the session: %v", err)
	}
	defer ctx.Session.Close()

	// hydrate cache on startup if configured
	if ctx.Config.HydrateCacheOnStart {
		if resp, err := ctx.Clients.Met.GetObjectIDs(); err == nil {
			log.Printf("successfully fetched %d object IDs from met api (cache will be hydrated)", resp.Total)
		} else {
			log.Fatalf("failed to hydrate cache with initial data: %v", err)
		}
	}

	// register commands
	log.Printf("registering %d bot command(s)\n", len(commands.Commands))
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
