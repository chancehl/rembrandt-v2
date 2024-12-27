package main

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
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
	var err error
	session, err = discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		log.Fatalf("invalid bot parameters: %v", err)
	}
}

func main() {
	session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	if err := session.Open(); err != nil {
		log.Fatalf("cannot open the session: %v", err)
	}

	defer session.Close()
}
