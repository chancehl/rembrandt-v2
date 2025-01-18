package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/chancehl/rembrandt-v2/internal/context"
	"github.com/chancehl/rembrandt-v2/internal/interactions"
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
func SubscribeCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate, ctx *context.BotContext) {
	channel, _ := utils.GetOption(i.Interaction, "channel")
	channelID := channel.ChannelValue(s).ID

	subscription, err := ctx.Clients.DB.GetSubscription(i.GuildID)
	if subscription != nil {
		interactions.RespondWithString(s, i, "Your guild already has an active subscription")
	}
	if err != nil {
		interactions.RespondWithDefaultErrorMessage(s, i)
	}

	_, err = ctx.Clients.DB.CreateSubscription(i.GuildID, channelID, i.Member.User.ID)
	if err != nil {
		interactions.RespondWithDefaultErrorMessage(s, i)
	}

	interactions.RespondWithString(s, i, fmt.Sprintf("I subscribed channel %s to daily updates", channel.ChannelValue(s).Name))
}
