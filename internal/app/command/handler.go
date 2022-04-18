package command

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func Handlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"match": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			// Log.
			event := fmt.Sprintf("id=%s channel_id=%s guild_id=%s guild_locale=%s ", i.ID, i.ChannelID, i.GuildID, i.GuildLocale)
			event += fmt.Sprintf("locale=%s user=%s nick=%s username=%s", i.Locale, i.User, i.Member.Nick, i.Member.User.Username)
			log.Println(event)

			// Access options in the order provided by the user.
			options := i.ApplicationCommandData().Options

			// Or convert the slice into a map
			optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
			for _, opt := range options {
				optionMap[opt.Name] = opt
			}

			// This example stores the provided arguments in an []interface{}
			// which will be used to format the bot's response
			margs := make([]interface{}, 0, len(options))
			msgformat := ""

			var match_id int
			if opt, ok := optionMap["match_id"]; ok {
				match_id = int(opt.IntValue())
				msgformat += "> Match %d\n"
				margs = append(margs, opt.IntValue())
			}

			var details bool
			if opt, ok := optionMap["details"]; ok {
				details = opt.BoolValue()
			}

			m, err := MatchResult(match_id)
			if err != nil {
				msgformat += "> Error: %s\n"
				margs = append(margs, err.Error())
			}

			if m.ID == "" {
				msgformat += "```text\n%s```\n"
				margs = append(margs, "Sorry, match not found.")
			}

			if err == nil && m.ID != "" {
				formatted := MatchFormat(m, details)
				msgformat += "```text\n%s```\n"
				margs = append(margs, formatted)
			}

			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				// Ignore type for now, they will be discussed in "responses"
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf(
						msgformat,
						margs...,
					),
				},
			})
			if err != nil {
				fmt.Println(err.Error())
			}
		},
		"demo": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			// Access options in the order provided by the user.
			options := i.ApplicationCommandData().Options

			// Or convert the slice into a map
			optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
			for _, opt := range options {
				optionMap[opt.Name] = opt
			}

			// This example stores the provided arguments in an []interface{}
			// which will be used to format the bot's response
			margs := make([]interface{}, 0, len(options))
			msgformat := "Demo\n"

			if opt, ok := optionMap["match_id"]; ok {
				margs = append(margs, opt.IntValue())
				msgformat += "> %d\n"
			}

			fmt.Printf("Margs: %s\n", margs)

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				// Ignore type for now, they will be discussed in "responses"
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf(
						msgformat,
						margs...,
					),
				},
			})
		},
	}

}
