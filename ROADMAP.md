# 6-Month Roadmap: Go Asset Service

This roadmap is the working plan for 24 weeks of Go learning and product development.

Goal: move from PHP developer learning Go to a backend profile that can credibly target roles requiring Go, APIs, SQL, Docker, Kubernetes, CI/CD, observability, cloud basics, and an AI proof of concept.

Current baseline (completed)
- Week 1 is already done in the codebase.
- Current app state: CLI app, JSON-backed asset data, validation, filtering, count by type, dynamic predicate filtering, unit tests.
- Current command flow lives in `cmd/cli/main.go` and domain logic is in `internal/assets`.

Working agreement
- All code and implementation changes in this project should be developed manually without AI code-generation assistance.
- AI should only be used for roadmap maintenance, planning, learning guidance, and answering specific questions when needed.
- Weekly progress should reflect what was built by hand, what was learned, and which questions required external guidance.

## How To Use This File Weekly

At the end of each week, update these items:
- Status: `not started`, `in progress`, or `done`
- Notes: what was finished, what was harder than expected, what should move to next week
- Evidence: PR, commit, screenshots, API examples, benchmarks, notes, or docs written
- Next adjustment: keep, cut, or split next week's scope

Suggested weekly time split
- 40% Go language and backend concepts
- 35% app development
- 15% tooling and DevOps
- 10% writing, English communication, and documentation

Weekly done rule
- Ship one visible improvement to the app
- Write at least one test for the new behavior
- Write a short weekly note in this file
- Keep the main branch runnable

---

## Month 1: Solid Go Foundations And Better CLI Design

### Week 1 - Baseline CLI And Validation
Status: done

Learning requirements
- Basic Go project layout
- Structs, methods, slices, maps, packages
- Flags, JSON loading, basic testing

App development requirements
- Load assets from JSON
- Filter assets
- Count assets by type
- Validate asset input
- Add tests for service and validation

Tools and practice
- `go test ./...`
- `go run cmd/cli/main.go`
- Git basics and small commits

Definition of done
- Completed in current repository state

### Week 2 - Refactor CLI Into Cleaner Layers
Status: done

Learning requirements
- Interfaces in Go
- Error handling style in Go
- Package boundaries and dependency direction
- Table-driven tests

App development requirements
- Separate CLI concerns from service concerns more clearly
- Introduce repository-style abstraction for asset storage
- Replace repetitive tests with table-driven tests where useful
- Improve CLI help and invalid-input messages

Tools and practice
- `gofmt ./...`
- `go test ./... -cover`
- Start using `golangci-lint`

Definition of done
- CLI parsing is thin
- Core asset logic is testable without the CLI
- At least one test file uses table-driven style

### Week 3 - CRUD In Memory Before Adding Persistence
Status: in progress

Learning requirements
- Pointers and value semantics
- Designing small services in Go
- Basic TDD loop for handlers and services

App development requirements
- Add create, update, delete asset operations in the service layer
- Add duplicate ID protection
- Add search by environment and status
- Extend tests for mutations and edge cases

Tools and practice
- Use subtests with `t.Run`
- Add make targets or task shortcuts for test and run flows

Definition of done
- Service supports full in-memory CRUD
- Mutation rules are covered by tests

### Week 4 - Persist Data To SQLite
Status: not started

Learning requirements
- SQL basics: SELECT, INSERT, UPDATE, DELETE, indexes
- Database access in Go with `database/sql`
- Differences between in-memory models and persisted data

App development requirements
- Replace JSON-only storage with SQLite persistence while keeping seed import support
- Add repository implementation for SQLite
- Create first database schema and seed script
- Add tests for repository behavior where practical

Tools and practice
- SQLite CLI or DB Browser for SQLite
- Migration tool such as `golang-migrate` or simple SQL migration files

Definition of done
- App can read and write assets from SQLite
- Schema is versioned in the repo

---

## Month 2: REST API, SQL, And Service Design

### Week 5 - Expose A REST API
Status: not started

Learning requirements
- `net/http`
- JSON request and response handling
- HTTP status codes and API design basics

App development requirements
- Create HTTP server entry point under `cmd/api`
- Add endpoints for list assets, get asset by ID, create asset
- Keep business logic in the service layer, not in handlers
- Add handler tests

Tools and practice
- Postman, Bruno, or curl collections
- `httptest`

Definition of done
- API serves the core read and create flows locally
- Handlers return correct status codes and validation errors

### Week 6 - Complete REST CRUD And Filtering
Status: not started

Learning requirements
- Query params, path params, request validation
- REST resource naming and pagination basics

App development requirements
- Add update and delete endpoints
- Support filtering by type, location, environment, and status via query params
- Add pagination basics: `limit` and `offset`
- Return consistent error payloads

Tools and practice
- Add API examples to README or dedicated docs
- Add integration-style handler tests

Definition of done
- Full CRUD works through HTTP
- Filtering is available through query params

### Week 7 - Better SQL Modeling And Query Quality
Status: not started

Learning requirements
- Normalization vs pragmatic denormalization
- Indexes and query plans
- Transactions and when to use them

App development requirements
- Review current asset schema and improve it where needed
- Add indexes for common queries
- Introduce transaction-safe update flow where appropriate
- Add repository benchmarks or basic performance checks

Tools and practice
- `EXPLAIN QUERY PLAN` in SQLite
- `go test -bench . ./...`

Definition of done
- Schema and queries match the main app access patterns
- At least one index is justified and documented

### Week 8 - Configuration And Environment Management
Status: not started

Learning requirements
- Config loading patterns in Go
- Environment variables, defaults, and secrets basics
- Twelve-factor app ideas

App development requirements
- Add config package for server port, DB path, log level, and environment
- Remove hard-coded runtime settings
- Add `.env.example` or config documentation
- Separate dev and test config flows

Tools and practice
- `direnv` or `.env` workflow
- Structured startup logging

Definition of done
- App starts from configuration rather than hard-coded values
- Local setup is documented clearly

---

## Month 3: Production Backend Basics

### Week 9 - Logging, Errors, And Middleware
Status: not started

Learning requirements
- Structured logging
- Error wrapping with `errors.Is` and `errors.As`
- Middleware patterns in Go HTTP servers

App development requirements
- Add request logging middleware
- Add panic recovery middleware
- Normalize service and handler error mapping
- Improve error messages for operators and API users

Tools and practice
- `log/slog` or `zerolog`
- Request IDs in logs

Definition of done
- Requests are logged consistently
- Errors are easier to trace across layers

### Week 10 - Authentication And Authorization Basics
Status: not started

Learning requirements
- API keys vs JWT
- Basic authn/authz concepts
- Protecting internal APIs pragmatically

App development requirements
- Add simple API key authentication for write endpoints
- Separate read and write permissions in a minimal way
- Test unauthorized and forbidden flows

Tools and practice
- Secrets via env vars
- API client examples with auth headers

Definition of done
- Write endpoints are protected
- Security behavior is tested

### Week 11 - Concurrency And Background Jobs
Status: not started

Learning requirements
- Goroutines, channels, contexts, cancellation
- Race conditions and the race detector
- Worker pattern basics

App development requirements
- Add a background import or sync job for asset data
- Support context cancellation in long-running operations
- Identify and remove unsafe shared-state patterns

Tools and practice
- `go test -race ./...`
- Context-aware logging

Definition of done
- At least one real use case uses goroutines safely
- Race detector passes for touched code

### Week 12 - Documentation And API Contract Quality
Status: not started

Learning requirements
- Writing technical documentation clearly in English
- API contracts and compatibility thinking
- Basic OpenAPI concepts

App development requirements
- Add OpenAPI spec for current endpoints
- Write architecture notes for packages and data flow
- Add request and response examples

Tools and practice
- Swagger or OpenAPI tooling
- Markdown docs in `docs/`

Definition of done
- A new developer can run and call the API from docs only
- API contract is documented and versioned

---

## Month 4: Containers, CI/CD, And Observability

### Week 13 - Containerize The Service
Status: not started

Learning requirements
- Dockerfile basics
- Multi-stage builds
- Image size and runtime concerns

App development requirements
- Add production-oriented Dockerfile
- Add local compose setup for API plus database
- Ensure config works in containers

Tools and practice
- Docker
- Docker Compose

Definition of done
- The service runs locally in containers with one command

### Week 14 - CI Pipeline With Quality Gates
Status: not started

Learning requirements
- CI fundamentals
- Build, test, lint, and artifact stages
- Failing fast and preserving developer feedback speed

App development requirements
- Add CI workflow for format, lint, test, and build
- Publish test coverage report or at least coverage output
- Make CI run on push and pull request

Tools and practice
- GitHub Actions
- `golangci-lint`

Definition of done
- Every change is validated automatically in CI
- Broken builds are visible immediately

### Week 15 - Metrics And Health Endpoints
Status: not started

Learning requirements
- RED metrics basics: rate, errors, duration
- Health vs readiness vs liveness
- Intro to Prometheus metrics

App development requirements
- Add `/healthz` and `/readyz`
- Add Prometheus metrics endpoint
- Track request count, latency, and error count

Tools and practice
- Prometheus
- Basic Grafana dashboard

Definition of done
- You can see request metrics and health state locally

### Week 16 - Kubernetes Basics
Status: not started

Learning requirements
- Pods, deployments, services, config maps, secrets
- Rolling updates and replica basics
- Why Kubernetes exists and what problems it solves

App development requirements
- Add Kubernetes manifests or Helm basics for the service
- Deploy locally to `kind` or `minikube`
- Configure app readiness and liveness probes

Tools and practice
- `kind` or `minikube`
- `kubectl`

Definition of done
- Service runs in local Kubernetes with working probes

---

## Month 5: Enterprise-Style Data And Platform Concerns

### Week 17 - CMDB Direction And Relationship Modeling
Status: not started

Learning requirements
- CMDB concepts: configuration items, ownership, dependencies
- Relationship modeling in relational databases
- Domain-driven naming basics

App development requirements
- Extend asset model with relationships such as `depends_on`, `owned_by`, or `runs_on`
- Add endpoints and queries for asset relationships
- Document the domain language clearly

Tools and practice
- ER diagrams
- SQL schema review notes

Definition of done
- App models more than flat assets; it captures relationships useful for a CMDB-style system

### Week 18 - ETL And Import Pipelines
Status: not started

Learning requirements
- ETL concepts: extract, transform, load
- Data cleansing and idempotent imports
- Batch processing tradeoffs

App development requirements
- Add import pipeline from CSV or external JSON feed into the database
- Track import job results and failures
- Add validation and deduplication rules for imports

Tools and practice
- CSV handling in Go
- Import logs and summary reports

Definition of done
- The app can ingest external asset data repeatably and safely

### Week 19 - Cloud Deployment Basics
Status: not started

Learning requirements
- Azure fundamentals or equivalent cloud basics
- Managed databases, app hosting, secrets, networking basics
- Cost awareness and least-privilege habits

App development requirements
- Deploy the service to a cloud sandbox
- Move config and secrets to the cloud platform model
- Document deployment architecture and tradeoffs

Tools and practice
- Azure App Service, Container Apps, or AKS equivalent
- Azure SQL or PostgreSQL equivalent if you outgrow SQLite

Definition of done
- The service is reachable in a cloud environment with documented setup

### Week 20 - Performance And Resilience
Status: not started

Learning requirements
- Profiling basics in Go
- Timeouts, retries, circuit-breaker thinking
- Load testing basics

App development requirements
- Add request timeouts and graceful shutdown
- Measure a few slow paths and improve one of them
- Add simple load test scenario for list and filter endpoints

Tools and practice
- `pprof`
- `hey` or `k6`

Definition of done
- You have at least one measured performance improvement and one resilience improvement

---

## Month 6: AI Proof Of Concept, Frontend Exposure, And Job Readiness

### Week 21 - AI-Assisted Asset Insights PoC
Status: not started

Learning requirements
- Practical AI system design for backend engineers
- Prompting, guardrails, and evaluation basics
- When to use AI vs normal code

App development requirements
- Build a small PoC that summarizes asset inventory risks or anomalies from your stored data
- Keep the AI integration behind a service interface
- Log prompts, outputs, and cost or rate constraints clearly

Tools and practice
- OpenAI-compatible API or Azure OpenAI equivalent
- Markdown evaluation notes with good and bad examples

Definition of done
- There is a small but real AI feature connected to asset data, not a fake demo

### Week 22 - Minimal Frontend And API Consumption
Status: not started

Learning requirements
- Enough Node.js and Angular or equivalent frontend knowledge to collaborate with frontend teams
- API consumption patterns and CORS basics

App development requirements
- Build a minimal frontend that lists assets, filters them, and shows relationships
- Keep scope small; this week is exposure, not frontend mastery
- Add one page for import job status or metrics view

Tools and practice
- Angular if you want direct job alignment, otherwise a small Node-backed UI plus notes on Angular concepts
- OpenAPI-generated client if practical

Definition of done
- You can demo the API through a small UI and explain the integration clearly

### Week 23 - Polish, Documentation, And Portfolio Quality
Status: not started

Learning requirements
- Explaining architecture decisions succinctly
- Writing README, runbooks, and tradeoff notes for interview discussion

App development requirements
- Clean the repository structure
- Improve README with architecture, setup, screenshots, API examples, and roadmap status
- Add ADR-style notes for major technical decisions

Tools and practice
- Simple architecture diagram
- Issue tracker or project board for completed milestones

Definition of done
- The repository reads like a serious backend portfolio project

### Week 24 - Interview Preparation And Gap Review
Status: not started

Learning requirements
- Review Go fundamentals, SQL, Docker, Kubernetes, observability, and cloud basics
- Practice explaining concurrency, interfaces, error handling, and system design in English

App development requirements
- Close the highest-value gaps discovered in review
- Fix rough edges, flaky tests, or missing docs
- Record a final project summary and lessons learned

Tools and practice
- Mock interview notes
- STAR stories from this project
- CV updates tied to shipped outcomes

Definition of done
- You can walk through the project end to end and defend the main engineering decisions

---

## Cross-Cutting Study Track

These run every week in parallel with the weekly feature work.

### Go Track
- Read a small amount every week from a high-quality Go source: official docs, blog posts, or a respected book
- Practice table-driven tests, interfaces, contexts, and error handling repeatedly
- Revisit one existing package each month and simplify it

### SQL Track
- Write SQL by hand every week
- Learn joins, indexes, transactions, and query plans through the app's real data model
- Move from SQLite to PostgreSQL once SQLite becomes a real limitation

### DevOps Track
- Keep CI green
- Containerize early and iterate
- Learn Kubernetes through one small deployment, not through theory only

### Observability Track
- Add logs first, then metrics, then dashboards
- Always ask: if this breaks in production, how would I know why?

### Communication Track
- Write one short technical note in English every week
- Practice explaining one design choice from the project in plain language
- Keep README and roadmap current

### AI Track
- Treat AI as an engineering feature, not a decoration
- Evaluate output quality against real asset-management use cases
- Prefer small, testable PoCs over broad claims

---

## Weekly Update Template

Copy this block under the current week when you review progress.

```md
Weekly review
- Status:
- What I finished:
- What I learned:
- What was difficult:
- Evidence:
- Scope moved to next week:
- Change for next week:
```

## End-of-Month Checkpoints

### End of Month 1
- Comfortable reading and writing small Go packages
- CLI and service layers are cleaner and easier to test
- SQLite introduced successfully

### End of Month 2
- A real REST API exists
- CRUD, filtering, config, and database basics are in place
- You can explain request flow from handler to DB

### End of Month 3
- App has logging, auth basics, concurrency exposure, and documentation
- You are thinking in production terms, not only in features

### End of Month 4
- The service is containerized, tested in CI, observable, and running on local Kubernetes

### End of Month 5
- The app looks more like a CMDB or federated asset service than a CRUD demo
- You have touched ETL, cloud deployment, and performance work

### End of Month 6
- The project is portfolio-ready
- You can demonstrate backend depth, tooling maturity, and practical curiosity about AI
- You have concrete evidence for salary and role discussions