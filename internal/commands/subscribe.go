package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/chancehl/rembrandt-v2/internal/clients/api/met"
	"github.com/chancehl/rembrandt-v2/internal/utils"
)

// `/subscribe` command definition
var SubscribeCommand = discordgo.ApplicationCommand{
	Name:                     "subscribe",
	Description:              "Susbcribes your discord for daily art updates",
	DefaultMemberPermissions: &manageServerPermission,
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "channel",
			Description: "The channel to subscribe to",
			Type:        discordgo.ApplicationCommandOptionChannel,
			Required:    true,
		},
	},
}

// `/subscribe` command definition
func SubscribeCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate, c *met.Client) {
	channel, _ := utils.GetOption(i.Interaction, "channel")

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("I should subscribe channel %s to daily updates", channel.ChannelValue(s).Name),
		},
	})
}
