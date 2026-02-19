---
name: migration-helper
description: Analyze GORM model changes, estimate resulting schema (DDL) differences, and propose safe migration steps with verification guidance.
license: MIT
metadata:
  mode: diagnostics
  purpose: migration
---

# Migration Helper

## When to Use

- A model change (field add/remove/type change) or migration code is proposed and you need a safe rollout plan.

## Rules

- Start with a non-destructive assessment and clearly call out destructive operations (drops, type narrows, index changes).
- Recommend backups, staging verification, and zero-downtime considerations for risky changes.

## Project-Specific Paths

- Core migration: `database/migrate/`
- Example app migration: `example2/internal/database/migrate/`
- Core models: `database/model/` (`Auth`, `TwoFA`, `TwoFABackup`, `TempEmail`)

## Workflow

1. Locate migration entrypoints in the paths above.
2. Map model edits to likely DDL changes and affected queries.
3. Recommend migration steps and verification commands (local + staging).

## Output

- **Affected models:** `path:line` entries.
- **Schema updates:** likely DDL changes and risk level (low/medium/high).
- **Verification:** steps and test commands.

## Examples

- "Adding a `Role` field to `Auth` model - what migration is needed?"
- "Changing `IDAuth` from uint64 to uint32 - what are the risks?"

## Related Skills

- `db-infra-mocks`, `test-runner`, `patch-applier`
