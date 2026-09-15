# Football Overlay Backend

Live-Fußball-Overlay mit Sportmonks-API Integration und league-spezifischen GUIs.

## Quickstart

```bash
cd backend
go run cmd/server/main.go
```

## API Endpoints

### Match Management
- `GET /healthz` - Health Check
- `GET/PUT /v1/matches/current` - Aktuelles Spiel
- `GET /v1/matches/current/overlay` - Standard HTML-Overlay
- `GET/PUT /v1/matches/current/config` - Overlay-Konfiguration
- `GET/PUT /v1/matches/{id}` - Match nach ID

### League GUIs
- `GET /v1/leagues` - Alle verfügbaren Ligen
- `GET /v1/leagues/{id}` - Spezifische Liga (bundesliga, champions-league, premier-league, dfb-pokal)
- `GET /v1/gui/{league}/overlay` - League-spezifisches Overlay

## Verfügbare League GUIs

### Bundesliga
```bash
http://localhost:8080/v1/gui/bundesliga/overlay
```
- Rot (#D20515), weiß
- DFL-Style
- Roboto Font

### Champions League
```bash
http://localhost:8080/v1/gui/champions-league/overlay
```
- Dunkelblau (#001f5c), Neon-Grün (#00ff85)
- Starball-Design
- Dunkles Theme

### Premier League
```bash
http://localhost:8080/v1/gui/premier-league/overlay
```
- Lila (#37003c), Neon-Grün (#00ff85)
- Löwe-Motiv
- Modernes Design

### DFB Pokal
```bash
http://localhost:8080/v1/gui/dfb-pokal/overlay
```
- Grün (#009639), weiß
- DFB-Style

## Testen

```bash
# Mock-Daten laden
curl -X PUT http://localhost:8080/v1/matches/bayern-dortmund-2025-04-12 \
  -H "Content-Type: application/json" \
  -d @examples/bayern-dortmund-2025.json

# Als current setzen
curl -X PUT http://localhost:8080/v1/matches/current \
  -H "Content-Type: application/json" \
  -d '{"matchId":"bayern-dortmund-2025-04-12"}'

# Overlay im Browser öffnen
# Standard:
http://localhost:8080/v1/matches/current/overlay

# Bundesliga GUI:
http://localhost:8080/v1/gui/bundesliga/overlay

# Champions League GUI:
http://localhost:8080/v1/gui/champions-league/overlay

# Premier League GUI:
http://localhost:8080/v1/gui/premier-league/overlay
```

## Docker

```bash
docker-compose up --build
```

## Android App

Siehe `../android-app/README.md` für Emulator-Setup und Overlay-Integration.
