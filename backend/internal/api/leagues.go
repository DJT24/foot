package api

import (
	"encoding/json"
	"net/http"
	"strings"
	"github.com/DJT24/foot/backend/internal/model"
)

// DetectLeague erkennt die Liga basierend auf Match-Daten
func DetectLeague(match *model.Match) *model.League {
	// Simple heuristic based on team names or external league ID
	// In production, this would come from the API response
	
	bundesligaTeams := []string{"Bayern", "Dortmund", "Leipzig", "Leverkusen", "Frankfurt", "Stuttgart"}
	
	for _, team := range bundesligaTeams {
		if strings.Contains(match.HomeTeam.Name, team) || strings.Contains(match.AwayTeam.Name, team) {
			return model.Leagues["bundesliga"]
		}
	}
	
	// Default to Bundesliga
	return model.Leagues["bundesliga"]
}

func (s *Server) handleLeagues(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Return all leagues
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(model.Leagues)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (s *Server) handleLeagueByID(w http.ResponseWriter, r *http.Request, leagueID string) {
	league, ok := model.Leagues[leagueID]
	if !ok {
		http.Error(w, "League not found", http.StatusNotFound)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(league)
}
