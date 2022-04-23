package si

type observer struct {
	Provider        provider        `json:"provider"`
	Map             mapp            `json:"map"`
	Round           round           `json:"round"`
	Player          player          `json:"player"`
	PhaseCountdowns phaseCountdowns `json:"phase_countdowns"`
	Grenades        grenades        `json:"grenades"`
	Previously      previously      `json:"previously"`
	// Auth auth `json:"auth"` // Maybe we dont need this here.
}

type provider struct {
	Name      string `json:"name"`      // "Counter-Strike: Global Offensive",
	Appid     int    `json:"appid"`     // 730,
	Version   int    `json:"version"`   // 13825,
	Steamid   string `json:"steamid"`   // "76561198342099054",
	Timestamp int    `json:"timestamp"` // 1650696739
}

type mapp struct {
	Mode  string `json:"mode"`  // "casual",
	Name  string `json:"name"`  // "de_dust2",
	Phase string `json:"phase"` // "live",
	Round int    `json:"round"` // 5,

	TeamCT team `json:"team_ct"`
	TeamT  team `json:"team_t"`

	NumMatchesToWinSeries int `json:"num_matches_to_win_series"` // 0,
	CurrentSpectators     int `json:"current_spectators"`        // 0,
	SouvenirsTotal        int `json:"souvenirs_total"`           // 0
}

type team struct {
	Score                  int `json:"score"`                    // 2,
	ConsecutiveRoundLosses int `json:"consecutive_round_losses"` // 1,
	TimeoutsRemaining      int `json:"timeouts_remaining"`       // 1,
	MatchesWonThisSeries   int `json:"matches_won_this_series"`  // 0
}

type round struct {
	Phase   string `json:"phase"`    // "live", "over"
	WinTeam string `json:"win_team"` // "CT"
}

type player struct {
	Steamid    string           `json:"steamid"`  // "76561198342099054",
	Name       string           `json:"name"`     // "D U S G",
	Team       string           `json:"team"`     // "T",
	Activity   string           `json:"activity"` // "playing",
	State      playerState      `json:"state"`
	Weapons    playerWeapons    `json:"weapons"`
	MatchStats playerMatchStats `json:"match_stats"`
}

type playerState struct {
	Health      int  `json:"health"`       // 100,
	Armor       int  `json:"armor"`        // 100,
	Helmet      bool `json:"helmet"`       // true,
	Flashed     int  `json:"flashed"`      // 0,
	Smoked      int  `json:"smoked"`       // 0,
	Burning     int  `json:"burning"`      // 0,
	Money       int  `json:"money"`        // 4650,
	RoundKills  int  `json:"round_kills"`  // 0,
	RoundKillhs int  `json:"round_killhs"` // 0,
	EquipValue  int  `json:"equip_value"`  // 4700
}

type playerWeapons struct {
	Weapons weapons `json:"weapons"`
}

type weapons struct {
	Weapon0 weapon `json:"weapon_0"`
	Weapon1 weapon `json:"weapon_1"`
	Weapon2 weapon `json:"weapon_2"`
	Weapon3 weapon `json:"weapon_3"`
	Weapon4 weapon `json:"weapon_4"`
	Weapon5 weapon `json:"weapon_5"`
}

type weapon struct {
	Name        string `json:"name"`          // "weapon_glock",
	Paintkit    string `json:"paintkit"`      // "hy_nerodia",
	Type        string `json:"type"`          // "Pistol",
	AmmoClip    int    `json:"ammo_clip"`     // 20,
	AmmoClipMax int    `json:"ammo_clip_max"` // 20,
	AmmoReserve int    `json:"ammo_reserve"`  // 120,
	State       string `json:"state"`         // "holstered"
}

type playerMatchStats struct {
	Kills   int `json:"kills"`   // 11,
	Assists int `json:"assists"` // 0,
	Deaths  int `json:"deaths"`  // 3,
	Mvps    int `json:"mvps"`    // 1,
	Score   int `json:"score"`   // 22
}

type phaseCountdowns struct {
	Phase       string `json:"phase"`         // "over",
	PhaseEndsIn string `json:"phase_ends_in"` // "2.3"
}

type grenades struct{}

type previously struct {
	Round  round  `json:"round"`
	Player player `json:"player"`
}
