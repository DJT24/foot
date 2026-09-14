#!/bin/bash
MATCH_ID="bayern-dortmund-2025-04-12"
MINUTE=$1
HOME_SCORE=$2
AWAY_SCORE=$3
curl -X PUT "http://localhost:8080/v1/matches/${MATCH_ID}" \
  -H "Content-Type: application/json" \
  -d "{\"id\":\"${MATCH_ID}\",\"homeTeam\":{\"name\":\"FC Bayern München\",\"shortName\":\"Bayern\",\"code\":\"FCB\"},\"awayTeam\":{\"name\":\"Borussia Dortmund\",\"shortName\":\"Dortmund\",\"code\":\"BVB\"},\"homeScore\":${HOME_SCORE},\"awayScore\":${AWAY_SCORE},\"minute\":${MINUTE},\"status\":\"LIVE\"}"
