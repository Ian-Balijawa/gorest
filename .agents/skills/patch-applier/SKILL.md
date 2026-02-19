---
name: patch-applier
description: Draft and apply small, atomic patches; always present a preview diff and appropriate verification commands.
license: MIT
metadata:
  mode: edit
  purpose: patch
---

# Patch Applier

## When to Use

- You have a focused code change that should be applied and reviewed as a single patch.

## Rules

- Keep patches atomic and single-purpose.
- Show a preview (paths + short hunk summary) before applying.
- After applying, run formatting, linters, and targeted tests as appropriate.

## Workflow

1. Locate edit points via `source-search` or `file-reader`.
2. Draft patch and show a preview for approval.
3. Apply patch and run `go fmt ./...` / `golangci-lint run ./...` / targeted `go test`.
4. Report modified files and verification results.

## Output

- **Patch summary:** 1-3 bullets.
- **Files changed:** list with `path:line`.
- **Verification:** commands and results.

## Related Skills

- `safe-edit-simulator`, `code-formatter`, `test-runner`
