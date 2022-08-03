package config

import (
	"errors"
	"log"
	"os"
	"path"

	"github.com/mitchellh/go-homedir"
)

type Config struct {
	Bot ConfigBot `yaml:"Bot"`
	API ConfigAPI `yaml:"API"`
}

type ConfigBot struct {
	// This is a Discord Token
	// Check the follow link for more:
	// https://discord.com/developers/docs/getting-started#adding-credentials
	Token string `yaml:"token"`
}

type ConfigAPI struct {
	Port int `yaml:"port"`
}

func loadConfig(file string) (config Config) {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}

	folder := path.Join(home, ".gcbot")
	_, err = os.Stat(folder)

	if errors.Is(err, os.ErrNotExist) {
		_, err = os.Create(folder)
		if err != nil {
			log.Fatal(err)
		}
	}

	// configFile := path.Join(folder, file)

	// cfg, err := ini.Load(path.Join(folder, file))
	// if err != nil {
	// 	log.Printf("Fail to read file: %v", err)
	// 	os.Exit(1)
	// }

	return
}

func NewBot() Config {
	return Config{}
}
