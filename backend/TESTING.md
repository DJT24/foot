# Testing Guide

## Lokale Tests

1. Backend starten:
   ```bash
   go run cmd/server/main.go
   ```

2. Test-Skript ausführen:
   ```bash
   ./test-overlay.sh
   ```

3. Score live updaten:
   ```bash
   ./update-score.sh 75 2 1
   ```

## Browser-Test

Overlay unter `http://localhost:8080/v1/matches/current/overlay` öffnen.

## OBS-Integration

1. YouTube-Video als Media Source
2. Browser-Source mit Overlay-URL darüber
