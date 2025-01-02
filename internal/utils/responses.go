package utils

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func RespondWithError(s *discordgo.Session, i *discordgo.InteractionCreate, err error) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Something went wrong: %v", err),
		},
	})
}
