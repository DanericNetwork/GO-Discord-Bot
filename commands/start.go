package commands

import (
	"github.com/bwmarrin/discordgo"
)

func Start(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "This is the start command, "+m.Author.Username)
}
