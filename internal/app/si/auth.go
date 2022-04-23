package si

type auth struct {
	Auth token `json:"auth"`
}

type token struct {
	Token string `json:"token"`
}
