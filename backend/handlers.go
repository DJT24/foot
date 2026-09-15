// Extend handlers.go to include league routes
// Add this to the existing ServeHTTP function:

/*
if strings.HasPrefix(path, "/v1/leagues/") {
	s.handleLeagues(w, r)
	return
}

if strings.HasPrefix(path, "/v1/gui/") {
	leagueID := strings.TrimPrefix(path, "/v1/gui/")
	s.handleGUIOverlay(w, r, leagueID)
	return
}
*/
