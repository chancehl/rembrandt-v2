package met

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func (o *Object) GenerateEmbed() *discordgo.MessageEmbed {
	fields := o.generateEmbedFields()

	embed := discordgo.MessageEmbed{
		Title:  o.Title,
		Fields: fields,
		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("The Metropolitan Museum of Art (Object ID: %d)", o.ObjectID),
		},
	}

	if o.PrimaryImage != "" {
		embed.Image = &discordgo.MessageEmbedImage{
			URL: o.PrimaryImage,
		}
	}

	return &embed
}

func (o *Object) generateEmbedFields() []*discordgo.MessageEmbedField {
	fields := []*discordgo.MessageEmbedField{}

	if o.ArtistDisplayName != "" {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "Artist",
			Value:  o.ArtistDisplayName,
			Inline: false,
		})
	}

	if o.Summary != "" {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "Description",
			Value:  o.Summary,
			Inline: false,
		})
	}

	if o.ObjectDate != "" {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "Date",
			Value:  o.ObjectDate,
			Inline: false,
		})
	}

	if o.Department != "" {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "Department",
			Value:  o.Department,
			Inline: false,
		})
	}

	if o.Culture != "" {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "Culture",
			Value:  o.Culture,
			Inline: false,
		})
	}

	if o.Period != "" {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "Period",
			Value:  o.Period,
			Inline: false,
		})
	}

	if o.Medium != "" {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "Medium",
			Value:  o.Medium,
			Inline: false,
		})
	}

	if o.AccessionNumber != "" {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "Accession Number",
			Value:  o.AccessionNumber,
			Inline: false,
		})
	}

	if o.ObjectURL != "" {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "URL",
			Value:  fmt.Sprintf("[View on The Met website](%s)", o.ObjectURL),
			Inline: false,
		})
	}

	return fields
}
