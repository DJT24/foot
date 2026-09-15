package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealthz(t *testing.T) {
	srv := NewServer()
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	w := httptest.NewRecorder()

	srv.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	if w.Body.String() != "ok" {
		t.Errorf("Expected body 'ok', got '%s'", w.Body.String())
	}
}

func TestGetCurrentMatch_NoMatch(t *testing.T) {
	srv := NewServer()
	req := httptest.NewRequest(http.MethodGet, "/v1/matches/current", nil)
	w := httptest.NewRecorder()

	srv.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", w.Code)
	}
}

func TestLeaguesEndpoint(t *testing.T) {
	srv := NewServer()
	req := httptest.NewRequest(http.MethodGet, "/v1/leagues", nil)
	w := httptest.NewRecorder()

	srv.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	if !strings.Contains(w.Body.String(), "bundesliga") {
		t.Error("Expected bundesliga in response")
	}
}

func TestGUIOverlayEndpoint(t *testing.T) {
	srv := NewServer()
	
	// Add a test match
	srv.matches["test-1"] = &Match{
		ID: "test-1",
		HomeTeam: Team{Name: "Home", ShortName: "HOM", Code: "HOM"},
		AwayTeam: Team{Name: "Away", ShortName: "AWA", Code: "AWA"},
		HomeScore: 1,
		AwayScore: 0,
		Minute: 45,
		Status: "LIVE",
	}
	srv.currentID = "test-1"

	req := httptest.NewRequest(http.MethodGet, "/v1/gui/bundesliga/overlay", nil)
	w := httptest.NewRecorder()

	srv.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	if !strings.Contains(w.Body.String(), "Bundesliga") {
		t.Error("Expected Bundesliga in overlay")
	}
}
