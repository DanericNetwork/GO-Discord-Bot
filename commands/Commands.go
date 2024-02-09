package commands

import (
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var Commands = map[string]func(s *discordgo.Session, m *discordgo.MessageCreate){
	"start": Start,
	"say":   Say,
	"insta": Insta,
}

func HandleCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	prefix := os.Getenv("PREFIX")
	if !strings.HasPrefix(m.Content, prefix) {
		return
	}

	args := strings.Split(m.Content, " ")
	input := args[0]

	command := strings.TrimPrefix(input, prefix)
	if handler, ok := Commands[command]; ok {
		handler(s, m)
	}
}
