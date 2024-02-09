package commands

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func Say(s *discordgo.Session, m *discordgo.MessageCreate) {
	args := strings.Split(m.Content, " ")

	if len(args) <= 1 {
		errorEmbed := &discordgo.MessageEmbed{
			Description: "Please provide a message to say",
			Color:       0xC20C00,
		}
		s.ChannelMessageSendEmbed(m.ChannelID, errorEmbed)
		return
	}

	message := strings.Join(args[1:], " ")
	letters := strings.Split(message, "")
	for i, letter := range letters {
		if !strings.ContainsAny(letter, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ") {
			continue
		}
		if letter != " " {
			letters[i] = ":regional_indicator_" + strings.ToLower(letter) + ":"
		}
	}
	newMessage := strings.Join(letters, "")
	// send embed
	embed := &discordgo.MessageEmbed{
		Description: newMessage,
		Color:       0x00B6C2,
	}
	s.ChannelMessageSendEmbed(m.ChannelID, embed)
}
