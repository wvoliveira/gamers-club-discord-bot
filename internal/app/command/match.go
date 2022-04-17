package command

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/elga-io/discord-bot-gamers-club/internal/pkg/co"
	"github.com/elga-io/discord-bot-gamers-club/internal/pkg/entity"
	"github.com/olekukonko/tablewriter"
)

// MatchResult get information from match.
func MatchResult(id int) (m entity.Match, err error) {
	resp, err := http.Get(fmt.Sprintf(co.Match, id))
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return
	}
	err = json.NewDecoder(resp.Body).Decode(&m)
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return
	}
	return
}

// MatchFormat return formated text from match info.
func MatchFormat(m entity.Match, details bool) (text string) {
	text = matchFormatResume(m)
	if details {
		text += "\n"
		text += matchFormatDetails(m)
	}
	return
}

// matchFormatResume render only basic information from match.
func matchFormatResume(m entity.Match) (text string) {
	data := [][]string{
		{m.TeamA, m.Games.ScoreA},
		{m.TeamB, m.Games.ScoreB},
		{"", ""},
		{"Data", m.Date},
		{"Duração", m.Games.Duration},
		{"Mapa", m.Games.MapName},
		{"Status", m.Status},
		//{"Demo", fmt.Sprintf(co.ObjectStore, m.Games.Demo)},
	}

	formated := &strings.Builder{}
	table := tablewriter.NewWriter(formated)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	//table.SetCenterSeparator("")
	//table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetBorder(false)
	table.SetTablePadding("  ")
	table.SetNoWhiteSpace(true)
	table.AppendBulk(data)
	table.Render()

	return formated.String()
}

// matchFormatDetails render details about match.
func matchFormatDetails(m entity.Match) (text string) {
	// Team A
	text += matchTeamTable(m.Games.Players.TeamA)
	text += fmt.Sprintf("Prob to win: %.2f%%\n", m.ProbWinA)
	text += "\n"

	// Team B
	text += matchTeamTable(m.Games.Players.TeamB)
	text += fmt.Sprintf("Prob to win: %.2f%%\n", m.ProbWinB)

	return text
}

// matchTeamTable format a table with team info.
func matchTeamTable(teams []entity.Team) (text string) {

	data := [][]string{}
	for _, team := range teams {
		data = append(data, []string{
			team.Player.Nick,
			team.Player.Level,
			team.NBKill,
			team.Assist,
			team.Death,
			team.Diff,
			team.ADR,
			team.KDR,
			team.Kast,
			team.Survived,
			team.Trade,
			team.FlashAssist,
			team.ClutchWon,
			team.MultiKills,
			team.FirstKill,
			team.RatingPoints,
		})
	}

	formated := &strings.Builder{}
	table := tablewriter.NewWriter(formated)
	table.SetHeader([]string{"Player", "Level", "K", "A", "D", "DIFF", "ADR", "KDR", "KAST", "S", "T", "FA", "1VSX", "MK", "FK", "RP"})
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetBorder(false)
	table.SetTablePadding("  ")
	table.SetNoWhiteSpace(true)
	table.AppendBulk(data)
	table.Render()
	return formated.String()
}
