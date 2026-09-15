package api

import (
	"net/http"
	"net/http/httptest"
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

func TestSetCurrentMatch(t *testing.T) {
	srv := NewServer()

	// First add a match
	matchJSON := `{"id":"test-1","homeTeam":{"name":"Home","shortName":"HOM","code":"HOM"},"awayTeam":{"name":"Away","shortName":"AWA","code":"AWA"},"homeScore":0,"awayScore":0,"minute":0,"status":"LIVE"}`
	req := httptest.NewRequest(http.MethodPut, "/v1/matches/test-1", nil)
	req.Body = nil
	srv.matches["test-1"] = &Match{
		ID: "test-1",
		HomeTeam: Team{Name: "Home", ShortName: "HOM", Code: "HOM"},
		AwayTeam: Team{Name: "Away", ShortName: "AWA", Code: "AWA"},
	}

	// Set as current
	setReq := httptest.NewRequest(http.MethodPut, "/v1/matches/current", nil)
	setReq.Body = nil
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, setReq)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
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
}
