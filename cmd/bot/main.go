package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/elga-io/discord-bot-gamers-club/internal/app/command"
	"github.com/elga-io/discord-bot-gamers-club/internal/app/si"
)

const (
	Version = "v0.0.1"
)

var (
	botToken string
	api      bool
	apiPort  string
	apiToken string
)

func init() {
	flag.StringVar(&botToken, "t", "", "Bot access token")
	flag.BoolVar(&api, "a", false, "Start a web API for receive actions from CS GO.")
	flag.StringVar(&apiPort, "ap", "8080", "API port. Default 8080")
	flag.StringVar(&apiToken, "at", "CCWJu64ZV3JHDT8hZc", "API token for authentication. Default: CCWJu64ZV3JHDT8hZc")
	flag.Parse()

	if botToken == "" {
		bt, ok := os.LookupEnv("DISCORD_BOT_GC_TOKEN")
		if ok {
			botToken = bt
		}
		if botToken == "" && !api {
			fmt.Println("You must set token for Discord API with -t flag or start API with -a flag.")
			os.Exit(3)
		}
	}
}

func main() {
	log.Println("Starting bot app..")

	if botToken != "" {
		go discord()
	}
	if api {
		si.API(apiPort, apiToken)
	}
}

func discord() {
	ds, err := discordgo.New("Bot " + botToken)
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
