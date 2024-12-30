package commands

import "github.com/bwmarrin/discordgo"

var UnsubscribeCommand = discordgo.ApplicationCommand{
	Name:                     "unsubscribe",
	Description:              "Unsubscribes your guild from daily art updates",
	DefaultMemberPermissions: &manageServerPermission,
}

func UnsubscribeCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Okay I will unsubscribe your guild from daily updates",
		},
	})
}
