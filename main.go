package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/chancehl/rembrandt-v2/commands"
	"github.com/chancehl/rembrandt-v2/config"
	"github.com/chancehl/rembrandt-v2/handlers"
	"github.com/joho/godotenv"
)

var botConfig *config.BotConfig

var session *discordgo.Session

var registrar *commands.SlashCommandRegistrar

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	botConfig = &config.BotConfig{
		TestGuildID:          os.Getenv("TEST_GUILD_ID"),
		RemoveCommandsOnExit: os.Getenv("REMOVE_COMMANDS_ON_EXIT") == "true",
	}
}

func init() {
	// create bot session
	var err error
	session, err = discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		log.Fatalf("invalid bot parameters: %v", err)
	}

	// register handlers
	session.AddHandler(handlers.OnBotReadyHandler)
	session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if handler, ok := handlers.CommandHandlers[i.ApplicationCommandData().Name]; ok {
			handler(s, i)
		}
	})

	// create command registrar
	registrar = commands.NewSlashCommandRegistrar(*botConfig, session, commands.CommandDefinitions)
}

func main() {
	log.Printf("starting bot with config %+v\n", *botConfig)

	// start the bot
	if err := session.Open(); err != nil {
		log.Fatalf("cannot open the session: %v", err)
	}
	defer session.Close()

	// register commands
	log.Printf("registering %d bot command(s)\n", len(commands.CommandDefinitions))
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
