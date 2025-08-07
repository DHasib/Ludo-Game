# Ludo Game

A web-based multiplayer Ludo game with a Go backend and React frontend.

## Project Structure

```
backend/   - Go API server and WebSocket hub
frontend/  - React client
migrations/ - SQL migrations
Dockerfiles and deployment configs provided
```

## Running Locally

1. Install Docker and Docker Compose.
2. Run `docker-compose up --build`.
3. Visit `http://localhost:3000` for the frontend.

## API

- `POST /api/auth/register` `{email, password}`
- `POST /api/auth/login` `{email, password}` -> `{token}`
- WebSocket endpoint: `ws://<host>/ws`

## Deployment

Kubernetes manifests in `deploy/k8s` provide basic deployments for backend, frontend, and PostgreSQL.

## Testing

Run unit tests:

```
cd backend
go test ./...
```

## Environment Variables

- `DATABASE_URL` – connection string for PostgreSQL.

## Notes

This is a scaffold demonstrating authentication, WebSocket setup, and deployment. Game logic and assets should be integrated as needed.
