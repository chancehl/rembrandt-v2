package commands

import "github.com/bwmarrin/discordgo"

var SubscribeCommand = discordgo.ApplicationCommand{
	Name:        "subscribe",
	Description: "Susbcribes your discord for daily art updates",
}
