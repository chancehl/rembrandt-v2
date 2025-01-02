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
	botConfig     *config.Config
	session       *discordgo.Session
	registrar     *commands.Registrar
	inMemoryCache *cache.InMemoryCache
	appContext    *context.AppContext
)

func init() {
	// load dot env variables
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	// create bot config
	botConfig = config.NewConfig(os.Getenv("TEST_GUILD_ID"), os.Getenv("REMOVE_COMMANDS_ON_EXIT") == "true")

	// create bot session
	var err error
	session, err = discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		log.Fatalf("invalid bot parameters: %v", err)
	}

	// create in memory cache
	inMemoryCache = cache.NewInMemoryCache()

	// create MET api client
	metClient := met.NewClient(inMemoryCache)

	// create DB client
	dbClient := db.NewClient()

	// create OpenAI client
	openAiClient := openai.NewClient()

	// create aggregated client struct
	clients := context.NewClientContext(metClient, dbClient, openAiClient)

	// create conetxt struct
	appContext = context.NewAppContext(clients, botConfig, session)

	// create command registrar
	registrar = commands.NewRegistrar(appContext)
}

func main() {
	log.Printf("starting bot with config %+v\n", *botConfig)

	// start the bot
	if err := session.Open(); err != nil {
		log.Fatalf("cannot open the session: %v", err)
	}
	defer session.Close()

	// cache objectIDs on startup
	log.Printf("hydrating cache with object IDs from met api")
	if objectIDsResponse, err := appContext.Clients.Met.GetObjectIDs(); err == nil {
		inMemoryCache.Set(met.ObjectIDsCacheKey, objectIDsResponse.ObjectIDs, met.ObjectIDsTTL)
		log.Printf("successfully hydrated cache with %d object IDs", objectIDsResponse.Total)
	} else {
		log.Fatalf("failed to hydrate cache with initial data: %v", err)
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
