package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/wvoliveira/gcbot/internal/app/command"
	"github.com/wvoliveira/gcbot/internal/pkg/config"
)

func main() {
	log.Println("Starting bot app..")
	cfg := config.New()
	discord(cfg.Token)
}

func discord(token string) {
	ds, err := discordgo.New(token)
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
