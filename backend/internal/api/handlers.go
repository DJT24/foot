package api

import (
	"encoding/json"
	"net/http"
	"strings"
	"sync"
)

type Server struct {
	mu          sync.RWMutex
	currentID   string
	matches     map[string]*Match
}

type Match struct {
	ID          string `json:"id"`
	HomeTeam    Team   `json:"homeTeam"`
	AwayTeam    Team   `json:"awayTeam"`
	HomeScore   int    `json:"homeScore"`
	AwayScore   int    `json:"awayScore"`
	Minute      int    `json:"minute"`
	Status      string `json:"status"`
}

type Team struct {
	Name      string `json:"name"`
	ShortName string `json:"shortName"`
	Code      string `json:"code"`
}

func NewServer() *Server {
	s := &Server{
		matches: make(map[string]*Match),
	}
	s.registerOverlayRoutes()
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if path == "/healthz" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
		return
	}

	if strings.HasPrefix(path, "/v1/matches/current") {
		s.handleCurrentMatches(w, r)
		return
	}

	if strings.HasPrefix(path, "/v1/matches/") {
		s.handleMatchesByID(w, r)
		return
	}

	http.NotFound(w, r)
}

func (s *Server) handleCurrentMatches(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.getCurrentMatch(w, r)
	case http.MethodPut:
		s.setCurrentMatch(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (s *Server) handleMatchesByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/v1/matches/")
	id = strings.TrimSuffix(id, "/overlay")
	id = strings.TrimSuffix(id, "/config")

	switch r.Method {
	case http.MethodGet:
		if strings.HasSuffix(r.URL.Path, "/overlay") {
			s.getOverlay(w, r, id)
		} else if strings.HasSuffix(r.URL.Path, "/config") {
			s.getConfig(w, r, id)
		} else {
			s.getMatchByID(w, r, id)
		}
	case http.MethodPut:
		if strings.HasSuffix(r.URL.Path, "/config") {
			s.updateConfig(w, r, id)
		} else {
			s.updateMatch(w, r, id)
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (s *Server) getCurrentMatch(w http.ResponseWriter, r *http.Request) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.currentID == "" {
		http.Error(w, "No current match set", http.StatusNotFound)
		return
	}

	match, ok := s.matches[s.currentID]
	if !ok {
		http.Error(w, "Current match not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(match)
}

func (s *Server) setCurrentMatch(w http.ResponseWriter, r *http.Request) {
	var req struct {
		MatchID string `json:"matchId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.currentID = req.MatchID
	w.WriteHeader(http.StatusOK)
}

func (s *Server) getMatchByID(w http.ResponseWriter, r *http.Request, id string) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	match, ok := s.matches[id]
	if !ok {
		http.Error(w, "Match not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(match)
}

func (s *Server) updateMatch(w http.ResponseWriter, r *http.Request, id string) {
	var match Match
	if err := json.NewDecoder(r.Body).Decode(&match); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.matches[id] = &match
	w.WriteHeader(http.StatusOK)
}
