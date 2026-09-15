package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/DJT24/foot/backend/internal/api"
)

type TestMatch struct {
	ID          string `json:"id"`
	HomeScore   int    `json:"homeScore"`
	AwayScore   int    `json:"awayScore"`
	Minute      int    `json:"minute"`
	Status      string `json:"status"`
}

func TestFullWorkflow(t *testing.T) {
	srv := api.NewServer()

	// 1. Load mock data
	mockData, err := os.ReadFile("../examples/bayern-dortmund-2025.json")
	if err != nil {
		t.Skipf("Mock data not available: %v", err)
	}

	var match TestMatch
	if err := json.Unmarshal(mockData, &match); err != nil {
		t.Fatalf("Failed to parse mock data: %v", err)
	}

	// 2. Set match
	req := httptest.NewRequest(http.MethodPut, "/v1/matches/bayern-dortmund-2025-04-12", nil)
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, req)

	// 3. Set as current
	setCurrentReq := httptest.NewRequest(http.MethodPut, "/v1/matches/current", nil)
	setCurrentW := httptest.NewRecorder()
	srv.ServeHTTP(setCurrentW, setCurrentReq)

	if setCurrentW.Code != http.StatusOK {
		t.Errorf("Expected status 200 for set current, got %d", setCurrentW.Code)
	}

	// 4. Get current match
	getCurrentReq := httptest.NewRequest(http.MethodGet, "/v1/matches/current", nil)
	getCurrentW := httptest.NewRecorder()
	srv.ServeHTTP(getCurrentW, getCurrentReq)

	if getCurrentW.Code != http.StatusOK {
		t.Errorf("Expected status 200 for get current, got %d", getCurrentW.Code)
	}

	// 5. Get overlay
	overlayReq := httptest.NewRequest(http.MethodGet, "/v1/matches/current/overlay", nil)
	overlayW := httptest.NewRecorder()
	srv.ServeHTTP(overlayW, overlayReq)

	if overlayW.Code != http.StatusOK {
		t.Errorf("Expected status 200 for overlay, got %d", overlayW.Code)
	}

	// 6. Get GUI overlay
	guiReq := httptest.NewRequest(http.MethodGet, "/v1/gui/bundesliga/overlay", nil)
	guiW := httptest.NewRecorder()
	srv.ServeHTTP(guiW, guiReq)

	if guiW.Code != http.StatusOK {
		t.Errorf("Expected status 200 for GUI overlay, got %d", guiW.Code)
	}

	t.Log("✅ Full workflow test passed")
}

func TestLeagueDetection(t *testing.T) {
	srv := api.NewServer()

	// Test Bundesliga detection
	leagueReq := httptest.NewRequest(http.MethodGet, "/v1/leagues/bundesliga", nil)
	leagueW := httptest.NewRecorder()
	srv.ServeHTTP(leagueW, leagueReq)

	if leagueW.Code != http.StatusOK {
		t.Errorf("Expected status 200 for league, got %d", leagueW.Code)
	}

	t.Log("✅ League detection test passed")
}
