package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/elga-io/discord-bot-gamers-club/internal/app/command"
)

var (
	BotToken = flag.String("t", "", "Bot access token")
)

var s *discordgo.Session

func init() {
	flag.Parse()

	if *BotToken == "" {
		bt, ok := os.LookupEnv("DISCORD_BOT_GC_TOKEN")
		if ok {
			*BotToken = bt
		}
		if *BotToken == "" {
			fmt.Println("You must set token for Discord API connection.")
			fmt.Println("With -t flag or DISCORD_BOT_GC_TOKEN env variable.")
			os.Exit(3)
		}
	}
}

func init() {
	var err error
	s, err = discordgo.New("Bot " + *BotToken)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
}

func init() {
	handlers := command.Handlers()
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := handlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}
