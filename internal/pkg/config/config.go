package config

import (
	"log"
	"path"

	"github.com/mitchellh/go-homedir"
)

var (
	homeFolder string

	configFiles []string = []string{
		path.Join(homeFolder, ".gcbot/gcbot.yaml"),
		"/etc/gcbot/gcbot.yaml",
	}

	defaultConfig = Config{
		ConfigBot{
			Token: "",
		},

		ConfigAPI{
			Port: 8080,
		},
	}
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

func init() {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}
	homeFolder = home
}

func LoadFile(file string) (config Config) {

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
