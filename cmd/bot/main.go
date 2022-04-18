package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/elga-io/discord-bot-gamers-club/internal/app/command"
)

var (
	BotToken string
)

func init() {
	flag.StringVar(&BotToken, "t", "", "Bot access token")
	flag.Parse()

	if BotToken == "" {
		bt, ok := os.LookupEnv("DISCORD_BOT_GC_TOKEN")
		if ok {
			BotToken = bt
		}
		if BotToken == "" {
			fmt.Println("You must set token for Discord API connection.")
			fmt.Println("With -t flag or DISCORD_BOT_GC_TOKEN env variable.")
			os.Exit(3)
		}
	}
}

func main() {
	ds, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}

	handlers := command.Handlers()
	ds.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := handlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	ds.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	err = ds.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	defer ds.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	log.Println("Gracefully shutting down.")
}
