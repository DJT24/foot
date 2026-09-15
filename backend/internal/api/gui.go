package api

import (
	"html/template"
	"net/http"
	"embed"
)

//go:embed templates/gui/*.html
var guiFS embed.FS

var guiTemplates = make(map[string]*template.Template)

func init() {
	// Load GUI templates
	tmpl, err := template.ParseFS(guiFS, "templates/gui/*.html")
	if err == nil {
		guiTemplates["bundesliga"] = tmpl
		guiTemplates["champions-league"] = tmpl
		guiTemplates["premier-league"] = tmpl
		guiTemplates["dfb-pokal"] = tmpl
	}
}

func (s *Server) handleGUIOverlay(w http.ResponseWriter, r *http.Request, leagueID string) {
	// Get current match
	s.mu.RLock()
	currentID := s.currentID
	match, ok := s.matches[currentID]
	s.mu.RUnlock()
	
	if !ok {
		http.Error(w, "No current match", http.StatusNotFound)
		return
	}
	
	// Get league config from model
	league := getLeagueConfig(leagueID)
	if league == nil {
		http.Error(w, "League not found", http.StatusNotFound)
		return
	}
	
	// Get overlay config
	configMu.RLock()
	cfg, ok := configStore[currentID]
	configMu.RUnlock()
	
	if !ok {
		cfg = &OverlayConfig{
			Opacity:  20,
			Blur:     10,
			Position: "top-center",
		}
	}
	
	// Render GUI template
	data := struct {
		Match  *Match
		League *LeagueConfig
		Config *OverlayConfig
	}{
		Match:  match,
		League: league,
		Config: cfg,
	}
	
	tmpl, ok := guiTemplates[leagueID]
	if !ok {
		http.Error(w, "Template not found", http.StatusNotFound)
		return
	}
	
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := tmpl.ExecuteTemplate(w, "gui.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// LeagueConfig is a local type for GUI templates
type LeagueConfig struct {
	ID             string
	Name           string
	ShortName      string
	Country        string
	LogoURL        string
	PrimaryColor   string
	SecondaryColor string
	FontFamily     string
	Theme          string
}

func getLeagueConfig(leagueID string) *LeagueConfig {
	leagues := map[string]*LeagueConfig{
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
			FontFamily:     "'Arial', sans-serif",
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
			FontFamily:     "'Arial', sans-serif",
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
	
	return leagues[leagueID]
}
