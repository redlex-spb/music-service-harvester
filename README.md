### Music Harvester

```bash
docker compose up --build
curl -X POST localhost:8080/v1/harvest/sources -H 'Content-Type: application/json' -d '{"urls":["https://open.spotify.com/track/123"]}'
```
