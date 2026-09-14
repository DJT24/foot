# Football Overlay Backend

Live-Fußball-Overlay mit Sportmonks-API Integration.

## Quickstart

```bash
cd backend
go run cmd/server/main.go
```

## API Endpoints

- `GET /healthz` - Health Check
- `GET/PUT /v1/matches/current` - Aktuelles Spiel
- `GET /v1/matches/current/overlay` - HTML-Overlay
- `GET/PUT /v1/matches/current/config` - Overlay-Konfiguration
- `GET/PUT /v1/matches/{id}` - Match nach ID

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
# http://localhost:8080/v1/matches/current/overlay
```

## Docker

```bash
docker-compose up --build
```
