package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/wvoliveira/discord-bot-gamers-club/internal/app/command"
	"github.com/wvoliveira/discord-bot-gamers-club/internal/app/si"
)

const (
	version = "v0.0.1"
)

var (
	botToken string
	api      bool
	apiPort  string
	apiToken string
)

func init() {
	flag.StringVar(&botToken, "t", "", "Bot access token")
	flag.Parse()

	if botToken == "" {
		bt, ok := os.LookupEnv("GCBOT_TOKEN")
		if ok {
			botToken = bt
		} else {
			print("You must set token for Discord BOT.")
			os.Exit(2)
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
