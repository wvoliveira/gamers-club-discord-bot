package entity

// Match root struct for match response.
// Example: https://gamersclub.com.br/lobby/match/15693476/1
type Match struct {
	Success      bool    `json:"success"`        // true
	Message      string  `json:"message"`        // null
	ID           string  `json:"id"`             // "15693476"
	Link         string  `json:"link"`           // "https://gamersclub.com.br/lobby/match/15693476"
	LinkFB       string  `json:"link_fb"`        // "https%3A%2F%2Fgamersclub.com.br%2Flobby%2Fmatch%2F15693476"
	Status       string  `json:"status"`         // "Finalizado"
	Date         string  `json:"data"`           // "15/04/2022 23:27"
	TeamA        string  `json:"time_a"`         // "Time adBrandini"
	TeamAAvatar  string  `json:"time_a_avatar"`  // "https://gcv1-assets.gamersclub.com.br/assets/images/playersBandeiras/br.png"
	TeamAIsElite bool    `json:"time_a_isElite"` // false
	AdminAvatarA string  `json:"admin_avatar_a"` // "https://static.gamersclub.com.br/players/avatar/1012987/1012987_medium.jpg"
	AdminAvatarB string  `json:"admin_avatar_b"` // "https://static.gamersclub.com.br/players/avatar/69494/69494_medium.jpg"
	TeamB        string  `json:"time_b"`         // "Time matheus"
	TeamBIsElite bool    `json:"time_b_isElite"` // false
	TeamBAvatar  string  `json:"time_b_avatar"`  // "https://gcv1-assets.gamersclub.com.br/assets/images/playersBandeiras/br.png"
	Format       string  `json:"formato"`        //"MD1"
	ProbWinB     float64 `json:"prob_win_b"`     // 50.62
	ProbWinA     float64 `json:"prob_win_a"`     // 49.38
}
