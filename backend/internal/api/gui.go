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
	// Load GUI templates for each league
	for leagueID := range []string{"bundesliga", "champions-league", "premier-league", "dfb-pokal"} {
		tmpl, err := template.ParseFS(guiFS, "templates/gui/*.html")
		if err == nil {
			guiTemplates[leagueID] = tmpl
		}
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
	
	// Get league config
	league, ok := model.Leagues[leagueID]
	if !ok {
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
		Match  *model.Match
		League *model.League
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
