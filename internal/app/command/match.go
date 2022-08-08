package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/wvoliveira/gcbot/internal/pkg/co"
	"github.com/wvoliveira/gcbot/internal/pkg/entity"
)

// MatchResult get information from match.
func MatchResult(id int) (m entity.Match, err error) {
	resp, err := http.Get(fmt.Sprintf(co.Match, id))
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return
	}

	err = json.Unmarshal(body, &m)
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return
	}

	if len(m.Games.RawPlayers) > 0 {
		switch m.Games.RawPlayers[0] {
		case '{':
			if err := json.Unmarshal(m.Games.RawPlayers, &m.Games.Players); err != nil {
				fmt.Printf("error: %s\n", err.Error())
			}
		}
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
	if len(m.Games.Players.TeamA) > 0 {
		players := m.Games.Players
		// Team A
		text += matchTeamTable(players.TeamA)
		text += fmt.Sprintf("Prob to win: %.2f%%\n", m.ProbWinA)
		text += "\n"

		// Team B
		text += matchTeamTable(players.TeamB)
		text += fmt.Sprintf("Prob to win: %.2f%%\n", m.ProbWinB)
	}

	if m.Status == "Live" {
		text += fmt.Sprintln("Jogo em andamento. Espere finalizar para coletar mais informações.")
	}

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
