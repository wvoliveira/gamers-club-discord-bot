package command

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// copied from discordgo.interactions struct
// useful to make the log explicit and get just we need.
// Tip: when GuildID is empty, the command was sent directly.
type event struct {
	ID        string
	GuildID   string
	ChannelID string
	User      string
	Locale    string
	Command   string
}

func LogEvent(i *discordgo.InteractionCreate) {
	e := event{
		ID:        i.ID,
		GuildID:   i.GuildID,
		ChannelID: i.ChannelID,
	}

	if i.User != nil {
		e.User = i.User.String()
	} else {
		e.User = i.Member.User.String()
	}

	e.Locale = i.Locale.String()

	// Use reflect to easy use in log.
	// TODO: test a library log that accept generic struct.
	v := reflect.ValueOf(e)
	typeOfS := v.Type()

	out := ""
	for i := 0; i < v.NumField(); i++ {
		key := strings.ToLower(typeOfS.Field(i).Name)
		value := v.Field(i).String()
		if value != "" {
			out += fmt.Sprintf("%s=%s ", key, value)
		}
	}

	out += fmt.Sprintf("command=%s", i.ApplicationCommandData().Name)
	log.Println(out)
}
