# Gopher Social

Gopher Social is a full-stack social platform for developers, built with Go (backend API) and React + Vite (frontend).

This project is designed for cloud deployment with:

- GCP (backend runtime)
- Vercel (frontend hosting)
- Supabase Postgres (database)
- Redis (caching)

## Features

### Core product features

- User registration with email invitation token
- Account activation via tokenized confirmation flow
- JWT-based authentication for protected endpoints
- Create, read, update, and soft-delete posts
- Follow and unfollow users
- Personalized feed model (including followed users)
- Post detail view with comment list
- Feed filters: pagination, sorting, tag filtering, and text search

### Authorization and security

- Role-based access model (`user`, `moderator`, `admin`) with precedence levels
- Ownership checks for post mutations
- JWT middleware with claim validation
- Basic Auth protection for runtime metrics endpoint
- Request body size limits and strict JSON field validation

### Performance and reliability

- Redis user caching layer (1-minute TTL)
- Fixed-window rate limiter middleware
- Database query timeouts and transactional write paths
- Optimistic locking for post updates (`version` field)
- Structured API response envelope for consistency

### Developer and ops tooling

- SQL migration workflow (up/down/create)
- Database seed generator for realistic test data
- Swagger/OpenAPI documentation endpoint
- Docker and Docker Compose for local dependencies
- Unit tests for middleware and API behavior

## Tech Stack

- Backend: Go, Chi router, JWT, Validator, Zap logging
- Frontend: React 18, TypeScript, Vite, SWR, React Router
- Database: PostgreSQL (Supabase)
- Cache: Redis
- Docs: Swagger
- Infra: Docker, Docker Compose, Vercel, GCP

## Local Development

### 1. Start dependencies

```bash
docker compose up -d db redis redis-commander
```

### 2. Run migrations

```bash
make migrate-up
```

### 3. (Optional) Seed sample data

```bash
make seed
```

### 4. Start API

```bash
air
```

### 5. Start frontend

```bash
cd web
npm install
npm run dev
```

Frontend default URL: `http://localhost:5173`

## Deployment Notes

### GCP

- Container image is built from the included `Dockerfile`.
- Runtime honors `PORT` (compatible with managed environments like Cloud Run).
- Ensure all app env vars are set in the service configuration.

### Vercel

- Deploy the `web/` app.
- Set `VITE_API_URL` to your API base URL.

### Supabase

- Use your Supabase Postgres connection string in `DB_ADDR`.
- Keep SSL enabled for hosted environments (`sslmode=require`).

### Redis

- Set `REDIS_ADDR`, `REDIS_PASSWORD`, and `REDIS_DB` according to your instance.
- Keep `REDIS_ENABLED=true` to enable cache-backed user lookups.

## Testing and Maintenance

Run tests:

```bash
make test
```

Generate Swagger docs:

```bash
make gen-docs
```

Create new migration:

```bash
make migration <name>
```

Rollback migration:

```bash
make migrate-down 1
```

## Frontend Scope

Current frontend includes:

- Login page
- Feed page
- Create post form
- Post detail page with comments
- Account confirmation page (`/confirm/:token`)

## Notes

- Mail provider integration exists for both Mailtrap and SendGrid; current API bootstrap path initializes Mailtrap client.
- Database schema includes comments, followers, invitations, roles, and indexed search support (GIN/pg_trgm).
