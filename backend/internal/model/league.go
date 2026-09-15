package model

type League struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	ShortName   string `json:"shortName"`
	Country     string `json:"country"`
	LogoURL     string `json:"logoURL"`
	PrimaryColor string `json:"primaryColor"`
	SecondaryColor string `json:"secondaryColor"`
	FontFamily  string `json:"fontFamily"`
	Theme       string `json:"theme"` // light, dark, custom
}

var Leagues = map[string]*League{
	"bundesliga": {
		ID:             "bundesliga",
		Name:           "Bundesliga",
		ShortName:      "BL",
		Country:        "Germany",
		LogoURL:        "https://cdn.sportmonks.com/images/soccer/leagues/18/82.png",
		PrimaryColor:   "#D20515",
		SecondaryColor: "#000000",
		FontFamily:     "'Roboto', sans-serif",
		Theme:          "light",
	},
	"champions-league": {
		ID:             "champions-league",
		Name:           "UEFA Champions League",
		ShortName:      "UCL",
		Country:        "Europe",
		LogoURL:        "https://cdn.sportmonks.com/images/soccer/leagues/5/5.png",
		PrimaryColor:   "#001f5c",
		SecondaryColor: "#00ff85",
		FontFamily:     "'UEFAChampionsLeague', sans-serif",
		Theme:          "dark",
	},
	"premier-league": {
		ID:             "premier-league",
		Name:           "Premier League",
		ShortName:      "PL",
		Country:        "England",
		LogoURL:        "https://cdn.sportmonks.com/images/soccer/leagues/19/83.png",
		PrimaryColor:   "#37003c",
		SecondaryColor: "#00ff85",
		FontFamily:     "'PremierLeague', sans-serif",
		Theme:          "dark",
	},
	"dfb-pokal": {
		ID:             "dfb-pokal",
		Name:           "DFB Pokal",
		ShortName:      "DFB",
		Country:        "Germany",
		LogoURL:        "https://cdn.sportmonks.com/images/soccer/leagues/13/109.png",
		PrimaryColor:   "#009639",
		SecondaryColor: "#FFFFFF",
		FontFamily:     "'Roboto', sans-serif",
		Theme:          "light",
	},
}
