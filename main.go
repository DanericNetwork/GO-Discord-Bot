package main

import (
	"fmt"
	"go-discord-bot/commands"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

const prefix string = "#"

type Command string

const (
	HelloCommand Command = "hello"
	SayCommand   Command = "say"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	token := os.Getenv("DISCORD_TOKEN")
	sess, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
	}

	sess.AddHandler(commands.HandleCommand)

	sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		args := strings.Split(m.Content, " ")
		if len(args) == 0 {
			return
		}

		input := args[0]
		command := Command(strings.TrimPrefix(input, prefix))

		switch command {
		case HelloCommand:
			if len(args) == 1 {
				s.ChannelMessageSend(m.ChannelID, "Hello, "+m.Author.Username)
				break
			}

			message := "Helloooooo, " + strings.Join(args[1:], " ")
			s.ChannelMessageSend(m.ChannelID, message)
		}
	})

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = sess.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	fmt.Println("Bot is running...")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
