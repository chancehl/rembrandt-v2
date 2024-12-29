package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/chancehl/rembrandt-v2/internal/parsers"
)

const ChannelOptionName = "channel"

var SubscribeCommand = discordgo.ApplicationCommand{
	Name:        "subscribe",
	Description: "Susbcribes your discord for daily art updates",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        ChannelOptionName,
			Description: "The channel to subscribe to",
			Type:        discordgo.ApplicationCommandOptionChannel,
			Required:    true,
		},
	},
}

func SubscribeCommandHandler(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	channelOption, _ := parsers.GetOption(interaction.Interaction, ChannelOptionName)

	session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Subscribed channel %v to daily art updates", channelOption.ChannelValue(session).Name),
		},
	})
}
