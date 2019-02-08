package auth

type Payload struct {
	Service string `json:"service"`
	Token   string `json:"token"`
}

type ServiceToken struct {
	Token string `json:"token"`
}
