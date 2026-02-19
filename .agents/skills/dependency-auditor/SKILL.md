---
name: dependency-auditor
description: Inspect Go module dependencies, detect outdated or vulnerable modules, and recommend safe updates or pinning strategies.
license: MIT
metadata:
  mode: analysis
  purpose: deps
---

# Dependency Auditor

## When to Use

- The user asks to audit `go.mod`/`go.sum` for outdated modules or known vulnerabilities.

## Responsibilities

- Run dependency analysis tools to identify updates and CVEs.
- Suggest minimal version bumps and `go.mod` edits, including tests to run after updates.

## Rules

- Do not modify `go.mod` without explicit approval.
- Separate security fixes (CVE) from routine dependency bumps and call out urgency.

## Commands

- `go list -m -u all` (list outdated modules)
- `govulncheck ./...` (check known vulnerabilities)
- `go mod tidy` (recommendation only, do not run without approval)

## Output

- Outdated modules with current and latest versions.
- Vulnerabilities (CVE) with severity and affected ranges.
- Recommended next steps and tests to run after updates.

## Related Skills

- `ci-orchestrator`, `static-analysis`
