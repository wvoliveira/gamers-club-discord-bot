package config

import (
	"flag"
	"log"
	"os"
)

var (
	defaultConfig = Config{
		Token: "",
	}
)

type Config struct {
	// This is a Discord Token
	// Check the follow link for more:
	// https://discord.com/developers/docs/getting-started#adding-credentials
	Token string
}

func (c *Config) loadEnv() {
	if bt, ok := os.LookupEnv("GCBOT_TOKEN"); ok {
		c.Token = bt
	}
}

func (c *Config) loadFlags() {
	flag.StringVar(&c.Token, "t", c.Token, "Bot access token")
	flag.Parse()
}

func New() (c Config) {
	c.loadEnv()
	c.loadFlags()

	if c.Token == "" {
		log.Println("Token is required. Use GCBOT_TOKEN or with -t flag.")
		os.Exit(2)
	}
	return
}
