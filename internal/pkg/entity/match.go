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
	Games        Games   `json:"jogos"`
}

type Games struct {
	ScoreA   string  `json:"score_a"`  // "14"
	ScoreB   string  `json:"score_b"`  // "16"
	MapName  string  `json:"map_name"` // "de_dust2"
	Demo     string  `json:"demo"`     //"2022-04-15__2332__1__15693476__de_dust2__timematheus__vs__timeadbrandini.zip"
	Players  Players `json:"players"`
	Duration string  `json:"duration"` // "42 Minutos"
}

type Players struct {
	TeamA []Team `json:"team_a"`
	TeamB []Team `json:"team_b"`
}

// Team in fact, this is information from player, not team.
// But in API response respect, I will keep the same name.
type Team struct {
	PlayerRoom              string `json:"player_room"`                // "a",
	IDBetaLobbyStatsPlayers string `json:"idbeta_lobby_stats_players"` // "276902295",
	IDLobbyGame             string `json:"idlobby_game"`               // "15693476",
	MatchID                 string `json:"match_id"`                   // null,
	IDPlayer                string `json:"idplayer"`                   // "978908",
	RoomID                  string `json:"room_id"`                    // "51716618",
	NBKill                  string `json:"nb_kill"`                    // "23",
	Assist                  string `json:"assist"`                     // "2",
	Death                   string `json:"death"`                      // "20",
	Point                   string `json:"point"`                      // "0",
	HS                      string `json:"hs"`                         // "10",
	Defuse                  string `json:"defuse"`                     // "0",
	Bombe                   string `json:"bombe"`                      // "0",
	TK                      string `json:"tk"`                         // "0",
	TKAssist                string `json:"tk_assist"`                  // "1",
	NB1                     string `json:"nb1"`                        // "0",
	NB2                     string `json:"nb2"`                        // "0",
	NB3                     string `json:"nb3"`                        // "0",
	NB4                     string `json:"nb4"`                        // "0",
	NB5                     string `json:"nb5"`                        // "0",
	NB1Kill                 string `json:"nb1kill"`                    // "10",
	NB2Kill                 string `json:"nb2kill"`                    // "5",
	NB3Kill                 string `json:"nb3kill"`                    // "1",
	NB4Kill                 string `json:"nb4kill"`                    // "0",
	NB5Kill                 string `json:"nb5kill"`                    // "0",
	PlusKill                string `json:"pluskill"`                   // "0",
	FirstKill               string `json:"firstkill"`                  // "8",
	Damage                  string `json:"damage"`                     // "3256",
	Hits                    string `json:"hits"`                       // "80",
	Shots                   string `json:"shots"`                      // "738",
	LastAlive               string `json:"last_alive"`                 // "0",
	ClutchWon               string `json:"clutch_won"`                 // "0",
	RoundsPlayed            string `json:"rounds_played"`              // "30",
	MapName                 string `json:"map_name"`                   // "de_dust2",
	MedalhaID               string `json:"medalhaID"`                  // "0",
	Level                   string `json:"level"`                      // "14",
	Rating                  string `json:"rating"`                     // "2114",
	Kast                    string `json:"kast"`                       // "21",
	Survived                string `json:"survived"`                   // "10",
	Trade                   string `json:"trade"`                      // "4",
	FlashAssist             string `json:"flash_assist"`               // "2",
	RatingPoints            string `json:"rating_points"`              // "-9",
	HitHeadshots            string `json:"hit_headshots"`              // "12",
	HitChest                string `json:"hit_chest"`                  // "36",
	HitStomach              string `json:"hit_stomach"`                // "22",
	HitLeftArm              string `json:"hit_left_arm"`               // "0",
	HitRightArm             string `json:"hit_right_arm"`              // "4",
	HitLeftLeg              string `json:"hit_left_leg"`               // "4",
	HitRightLeg             string `json:"hit_right_leg"`              // "2",
	CreatedAt               string `json:"created_at"`                 // "2022-04-16 00:14:51",
	UpdatedAt               string `json:"updated_at"`                 // null,
	DiffPositivo            bool   `json:"diffPositivo"`               // true,
	Diff                    string `json:"diff"`                       // "3",
	KDR                     string `json:"kdr"`                        // "1.15",
	Phs                     string `json:"phs"`                        // "43",
	ADR                     string `json:"adr"`                        // "108.53",
	DiffNegativo            bool   `json:"diffNegativo"`               // true,
	Pkast                   string `json:"pkast"`                      // "70",
	MultiKills              string `json:"multikills"`                 // "6",
	PlAvatar                string `json:"plAvatar"`                   // "players/avatar/978908/978908",
	Banido                  bool   `json:"banido"`                     // false,
	CompleteFor             string `json:"complete_for"`               // "0",
	CompleteWho             string `json:"complete_who"`               // "0",
	IsDoubleRatingPoints    string `json:"isDoubleRatingPoints"`       // "0",
	IsFreezeRatingPoints    string `json:"isFreezeRatingPoints"`       // "0",
	IsVaultDrop             bool   `json:"isVaultDrop"`                // false,
	VaultSlug               string `json:"vaultSlug"`                  // null,
	VaultImage              string `json:"vaultImage"`                 // null,
	Player                  Player `json:"player"`
	ZeroRP                  bool   `json:"zeroRP"`     // false,
	PositiveRP              bool   `json:"positiveRP"` // false,
	DoubleRP                bool   `json:"doubleRP"`   // false,
	FreezeRP                bool   `json:"freezeRP"`   // false
}

type Player struct {
	ID                string `json:"id"`                // "978908",
	Reputation        string `json:"reputation"`        // "100",
	Avatar            string `json:"avatar"`            // "players/avatar/978908/978908",
	AvatarName        string `json:"avatarName"`        // null,
	AvatarExtension   string `json:"avatarExtension"`   // "jpg",
	Nationality       string `json:"nationality"`       // "BR",
	Banned            bool   `json:"banned"`            // false,
	Nick              string `json:"nick"`              // "TLHXfps",
	Level             string `json:"level"`             // "14",
	FeaturedMedal     string `json:"featuredMedal"`     // null,
	IsPremium         string `json:"isPremium"`         // null,
	IsPlus            string `json:"isPlus"`            // null,
	IsArgentina       string `json:"isArgentina"`       // null,
	IsElite           string `json:"isElite"`           // null,
	IsSubscriber      bool   `json:"isSubscriber"`      // false,
	SubscriptionMedal int    `json:"subscriptionMedal"` // null,
	Type              string `json:"type"`              // "free",
	AvatarHtml        string `json:"avatarHtml"`        // "<div class=\"gc-avatar medium\">\n    <span class=\"gc-avatar-frame\" style=\"background-image: url();\"></span>\n    <svg viewBox=\"0 0 42 42\" class=\"gc-avatar-donut\">\n        <circle class=\"donut-hole\" cx=\"21\" cy=\"21\" r=\"15.91549430918954\" fill=\"transparent\"></circle>\n        <circle class=\"donut-ring\" cx=\"21\" cy=\"21\" r=\"15.91549430918954\" fill=\"transparent\" stroke=\"#92AA00\"></circle>\n        <circle class=\"donut-segment\" cx=\"21\" cy=\"21\" r=\"15.91549430918954\" fill=\"transparent\" stroke=\"#000\" stroke-opacity=\".7\" stroke-width=\"1\" stroke-dasharray=\"0 100\" stroke-dashoffset=\"25\"></circle>\n    </svg>\n    <img\n        class=\"gc-avatar-image\"\n        src=\"https://static.gamersclub.com.br/players/avatar/978908/978908_medium.jpg\"\n        alt=\"TLHXfps\"\n        onerror=\"this.style.display = 'none';this.setAttribute('error', true);\" />\n    <span class=\"gc-avatar-fallback\"></span>\n</div>\n",
	LevelHtml         string `json:"levelHtml"`         // "\n  <span class=\"badge-level badge-level-14 \">\n    14  </span>\n\n"
}
