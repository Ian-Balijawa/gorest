---
name: code-formatter
description: Apply repository-standard formatting to minimize diff noise and meet linter expectations.
license: MIT
metadata:
  mode: edit
  purpose: format
---

# Code Formatter

## When to Use

- After making code edits and before running linters/tests or opening a PR.

## Rules

- Prefer `go fmt ./...` for Go files; do not perform sweeping reformatting unless requested.
- Only format files you changed unless the user asks for a global reformat.
- Run `gofmt`/`goimports` where appropriate and include updated imports in the patch.

## Commands

- **Format:** `go fmt ./...` (or `gofmt -w .` when explicitly requested)
- **Verify:** `golangci-lint run --timeout 5m --verbose ./...` after formatting

## Output

- **Files formatted:** list of changed files.
- **Next:** recommended lint or test command.

## Related Skills

- `patch-applier`, `linter-runner`
