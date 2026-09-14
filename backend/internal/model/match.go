package model

type Match struct {
	ID        string `json:"id"`
	HomeTeam  Team   `json:"homeTeam"`
	AwayTeam  Team   `json:"awayTeam"`
	HomeScore int    `json:"homeScore"`
	AwayScore int    `json:"awayScore"`
	Minute    int    `json:"minute"`
	Status    string `json:"status"`
}

type Team struct {
	Name      string `json:"name"`
	ShortName string `json:"shortName"`
	Code      string `json:"code"`
}
