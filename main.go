package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/chancehl/rembrandt-v2/config"
	"github.com/chancehl/rembrandt-v2/handlers"
	"github.com/joho/godotenv"
)

var botConfig *config.BotConfig

var session *discordgo.Session

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
	// create session
	var err error
	session, err = discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		log.Fatalf("invalid bot parameters: %v", err)
	}

	// register event handlers
	session.AddHandler(handlers.HandleOnBotReadyEvent)

	// register slash commands
	// session.AddHandler(commands.ArtCommand)
}

func main() {
	if err := session.Open(); err != nil {
		log.Fatalf("cannot open the session: %v", err)
	}
	defer session.Close()

	log.Printf("starting bot with config %+v\n", *botConfig)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("bot has started (press ctrl+c to exit)")
	<-stop

	if botConfig.RemoveCommandsOnExit {
		log.Println("removing bot commands...")
	}
	log.Println("bot exited gracefully")
}
