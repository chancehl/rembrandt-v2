package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/chancehl/rembrandt-v2/internal/context"
	"github.com/chancehl/rembrandt-v2/internal/interactions"
)

var UnsubscribeCommand = discordgo.ApplicationCommand{
	Name:                     "unsubscribe",
	Description:              "Unsubscribes your guild from daily art updates",
	DefaultMemberPermissions: &manageServerPermission,
}

func UnsubscribeCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate, ctx *context.BotContext) {
	subscription, err := ctx.Clients.DB.GetSubscription(i.GuildID)
	if err != nil {
		interactions.RespondWithDefaultErrorMessage(s, i)
	}
	if subscription == nil {
		interactions.RespondWithString(s, i, "Your server does not have an active subscription")
	}
	if err := ctx.Clients.DB.ActivateSubscription(i.GuildID, i.User.ID); err != nil {
		log.Printf("failed to deactivate guild %s subscription: %v", i.GuildID, err)
		interactions.RespondWithDefaultErrorMessage(s, i)
	}
	interactions.RespondWithString(s, i, "Okay, I will unsubscribe your server from daily updates.")
}
