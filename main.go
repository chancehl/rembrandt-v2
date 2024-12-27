package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/chancehl/rembrandt-v2/handlers"
	"github.com/joho/godotenv"
)

var (
	GuildID        = "927013651339157575"
	RemoveCommands = true
)

var session *discordgo.Session

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env file: %v", err)
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
	session.AddHandler(handlers.OnBotReady)

	// register slash commands
	// session.AddHandler(commands.ArtCommand)
}

func main() {
	if err := session.Open(); err != nil {
		log.Fatalf("cannot open the session: %v", err)
	}
	defer session.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("bot has started (press ctrl+c to exit)")
	<-stop

	if RemoveCommands {
		log.Println("removing bot commands...")
	}
	log.Println("bot exited gracefully")
}
