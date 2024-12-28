package commands

import "github.com/bwmarrin/discordgo"

var SubscribeCommandDefinition = discordgo.ApplicationCommand{
	Name:        "subscribe",
	Description: "Susbcribes your discord for daily art updates",
}
