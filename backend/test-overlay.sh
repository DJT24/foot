#!/bin/bash
set -e
echo "🚀 Starte Backend-Test..."
curl -s http://localhost:8080/healthz > /dev/null && echo "✅ Backend läuft" || (echo "❌ Backend nicht erreichbar" && exit 1)
echo "📦 Lade Mock-Daten..."
curl -X PUT http://localhost:8080/v1/matches/bayern-dortmund-2025-04-12 -H "Content-Type: application/json" -d @examples/bayern-dortmund-2025.json > /dev/null && echo "✅ Mock-Daten geladen"
echo "⚽ Setze als aktuelles Spiel..."
curl -X PUT http://localhost:8080/v1/matches/current -H "Content-Type: application/json" -d '{"matchId":"bayern-dortmund-2025-04-12"}' > /dev/null && echo "✅ Aktuelles Spiel gesetzt"
echo "🎨 Öffne Overlay im Browser..."
open http://localhost:8080/v1/matches/current/overlay 2>/dev/null || echo "📺 Overlay URL: http://localhost:8080/v1/matches/current/overlay"
echo "✅ Test abgeschlossen!"
