package api

import (
	"embed"
	"encoding/json"
	"html/template"
	"net/http"
	"sync"
)

//go:embed templates/overlay.html
var embedFS embed.FS

var overlayTemplate = template.Must(template.New("overlay").ParseFS(embedFS, "templates/overlay.html"))

type OverlayConfig struct {
	Opacity       int      `json:"opacity"`
	Blur          int      `json:"blur"`
	Position      string   `json:"position"`
	ShowScorers   bool     `json:"showScorers"`
	ShowCards     bool     `json:"showCards"`
	ShowStats     bool     `json:"showStats"`
	VisibleFields []string `json:"visibleFields"`
}

var (
	configStore = make(map[string]*OverlayConfig)
	configMu    sync.RWMutex
)

func (s *Server) registerOverlayRoutes() {
	// Routes werden über ServeHTTP gehandled
}

func (s *Server) getOverlay(w http.ResponseWriter, r *http.Request, matchID string) {
	s.mu.RLock()
	match, ok := s.matches[matchID]
	s.mu.RUnlock()

	if !ok {
		http.Error(w, "Match not found", http.StatusNotFound)
		return
	}

	configMu.RLock()
	cfg, ok := configStore[matchID]
	configMu.RUnlock()

	if !ok {
		cfg = &OverlayConfig{
			Opacity:       20,
			Blur:          10,
			Position:      "top-center",
			ShowScorers:   true,
			ShowCards:     true,
			ShowStats:     false,
			VisibleFields: []string{"score", "time", "teams"},
		}
	}

	data := struct {
		Match  *Match
		Config *OverlayConfig
	}{
		Match:  match,
		Config: cfg,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := overlayTemplate.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *Server) getConfig(w http.ResponseWriter, r *http.Request, matchID string) {
	configMu.RLock()
	defer configMu.RUnlock()

	cfg, ok := configStore[matchID]
	if !ok {
		cfg = &OverlayConfig{
			Opacity:       20,
			Blur:          10,
			Position:      "top-center",
			ShowScorers:   true,
			ShowCards:     true,
			ShowStats:     false,
			VisibleFields: []string{"score", "time", "teams"},
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cfg)
}

func (s *Server) updateConfig(w http.ResponseWriter, r *http.Request, matchID string) {
	var cfg OverlayConfig
	if err := json.NewDecoder(r.Body).Decode(&cfg); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	configMu.Lock()
	defer configMu.Unlock()

	configStore[matchID] = &cfg
	w.WriteHeader(http.StatusOK)
}
