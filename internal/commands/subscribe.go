package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/chancehl/rembrandt-v2/internal/models"
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

func SubscribeCommandHandler(session *models.BotSession, interaction *models.BotInteraction) {
	channelOption, _ := interaction.GetOption(ChannelOptionName)
	content := fmt.Sprintf("Subscribed channel %v to daily art updates", channelOption.ChannelValue(session.Session).Name)
	session.RespondToInteractionWithString(interaction, content)
}
