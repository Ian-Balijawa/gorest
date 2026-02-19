---
name: db-infra-mocks
description: Propose minimal seams and local substitutes so tests run without real RDBMS/Redis/Mongo infrastructure.
license: MIT
metadata:
  mode: design
  purpose: infra
---

# DB Infrastructure Mocks

## When to Use

- Integration dependencies (RDBMS, Redis, MongoDB) prevent running tests or local development and you need safe, testable alternatives.

## Responsibilities

- Locate where clients and connections are constructed and consumed (`database/` package).
- Propose interface-based seams and minimal mock implementations.
- Provide migration path and verification steps so production behavior remains unchanged.

## Rules

- Prefer small interfaces and constructor injection over large refactors.
- Make tests deterministic and avoid network calls.
- Provide unit and integration verification steps for any change.

## Project-Specific Details

- RDBMS connections: `gdb.InitDB()` / `gdb.GetDB()` (GORM `*gorm.DB`).
- Redis connections: `gdb.InitRedis()` / `gdb.GetRedis()` (radix v4 `radix.Client`).
- MongoDB connections: `gdb.InitMongo()` / `gdb.GetMongo()` (`*mongo.Client` from official driver v2).
- Close all: `gdb.CloseAllDB()`.

## Recommended Approaches

- **RDBMS:** Use SQLite in-memory (`driver: "sqlite3"`) for simple GORM-backed tests.
- **Redis:** Use a radix v4 compatible mock or inject an interface wrapper around `radix.Client`.
- **MongoDB:** Inject a mock repository interface or use a local test MongoDB instance.

## Output

- Dependency map with `path:line` of client construction and usage.
- Recommended seam (interface signature) and where to inject it.
- Example mock/test code snippet and verification commands.

## Related Skills

- `test-runner`, `patch-applier`, `migration-helper`
