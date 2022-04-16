package command

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/elga-io/discord-bot-gamers-club/internal/pkg/entity"
)

// MatchResult get information from match.
// Ex.: https://gamersclub.com.br/lobby/match/15693476/1
func MatchResult(id int) (m entity.Match, err error) {
	resp, err := http.Get(fmt.Sprintf("https://gamersclub.com.br/lobby/match/%d/1", id))
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return
	}
	err = json.NewDecoder(resp.Body).Decode(&m)
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return
	}

	fmt.Printf("Status: %s\n", m.Status)
	return
}
