package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/chancehl/rembrandt-v2/internal/helpers/interaction"
	"github.com/chancehl/rembrandt-v2/internal/helpers/responses"
)

var SubscribeCommand = discordgo.ApplicationCommand{
	Name:        "subscribe",
	Description: "Susbcribes your discord for daily art updates",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "channel",
			Description: "The channel to subscribe to",
			Type:        discordgo.ApplicationCommandOptionChannel,
			Required:    true,
		},
	},
}

func SubscribeCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	channel, _ := interaction.GetOption(i.Interaction, "channel")
	content := fmt.Sprintf("I should subscribe channel %s to daily updates", channel.ChannelValue(s).Name)
	responses.RespondWithString(s, i.Interaction, content)
}
