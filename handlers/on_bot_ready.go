package handlers

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func OnBotReadyHandler(s *discordgo.Session, r *discordgo.Ready) {
	log.Printf("logged in as: %s#%s", s.State.User.Username, s.State.User.Discriminator)
}
