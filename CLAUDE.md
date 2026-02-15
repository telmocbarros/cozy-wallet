# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Cozy Wallet is a digital wallet backend service focused exclusively on **payment cards** (debit and credit). Loyalty cards, boarding passes, tickets, and other wallet item types are out of scope.

## Technology Stack

- **Language**: Go (latest stable)
- **Database**: PostgreSQL (raw SQL via `database/sql` or `pgx` — no ORM)
- **Router**: `net/http` with `chi`
- **Migrations**: `golang-migrate`
- **Config**: Environment variables
- **Testing**: Go's built-in `testing` package
- **Auth**: bcrypt for passwords, JWT for access tokens

## Build & Development Commands

```bash
go build ./...
go test ./...
go test ./path/to/package -run TestName
go test -cover ./...
go vet ./...
```

## Architecture

Clean layered architecture with dependencies flowing inward: `handler → service → repository`. Never skip layers.

```
cmd/api/main.go              # Entrypoint
internal/
  config/                    # Configuration loading
  database/                  # DB connection and migration setup
  handler/                   # HTTP handlers (transport layer)
  middleware/                 # HTTP middleware (logging, auth, recovery)
  model/                     # Domain types and structs (pure data, no side effects)
  repository/                # Database access layer (SQL queries, row mapping)
  service/                   # Business logic and orchestration
  validator/                 # Input validation helpers
migrations/                  # SQL migration files
```

### Layer Rules

- **handler**: Parses HTTP requests, calls service, writes HTTP responses. No business logic. Maps domain errors to HTTP responses here only.
- **service**: All business rules, validation, authorization, domain logic. Calls repository. Returns domain errors, not HTTP statuses.
- **repository**: Executes SQL, maps rows to models. No business logic, no HTTP awareness.
- **model**: Shared pure data structures across layers.

## API Conventions

### Response Envelope

All endpoints use a consistent JSON envelope:

- Success: `{"status": "success", "data": {...}}`
- Success with pagination: `{"status": "success", "data": [...], "pagination": {"page": 1, "per_page": 20, "total": 58}}`
- Error: `{"status": "error", "error": {"code": "VALIDATION_ERROR", "message": "..."}}`

### HTTP Status Codes

`201` creation, `200` success, `204` deletion, `400` bad input, `404` not found, `409` conflict, `500` unexpected.

### Error Handling

Use custom application error types carrying an error code, user-safe message, and HTTP status. Log internal details server-side; never leak stack traces to the client.

## Implementation Phases

### Phase 1: Card CRUD

`POST/GET/PUT/DELETE /api/v1/cards[/{id}]` — Store only safe metadata (last four digits, network, cardholder name, expiry, card type). Never store full card numbers. Soft-delete via `deleted_at`. All list endpoints paginated.

### Phase 2: Transaction Operations

`POST/GET /api/v1/cards/{id}/transactions[/{txn_id}]` plus search endpoint. Transaction types: `debit`, `credit`, `refund`. Statuses: `pending`, `completed`, `failed`. Transactions are immutable — corrections via offsetting transactions. Compute running balance per card.

### Phase 3: Authentication

`POST /api/v1/auth/register` and `/login`. bcrypt passwords, short-lived JWT access tokens. Auth middleware injects user ID into context. Retrofit all endpoints with ownership checks via `user_id` foreign key on cards.

## Key Rules

- Write SQL directly — no ORM
- Handle errors explicitly; never silently swallow them
- Use transactions where data consistency matters
- Keep handlers stateless; use connection pooling
- Paginate all list endpoints
- Every public function in service/repository layers should be testable without a full HTTP server
- Do not add features or abstractions beyond the current phase