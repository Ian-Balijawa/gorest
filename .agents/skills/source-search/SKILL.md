---
name: source-search
description: Fast, repository-wide search to locate symbols, routes, env vars, error codes, and other artifacts (read-only).
license: MIT
metadata:
  mode: read-only
  purpose: search
---

# Source Search

## When to Use

- You need to locate where a symbol, string, env var, or error code is defined or referenced across the codebase.

## Rules

- Read-only: do not modify files.
- Prefer `grep` for broad discovery, then targeted `read` for the best matches.
- Narrow with glob patterns to reduce noise.

## Workflow

1. Identify exact search keys and likely variants.
2. Glob by folder or pattern to limit scope.
3. Grep and rank results by relevance.
4. Return top matches with `path:line` and a short note.

## Output

- Definitions and top usages with `path:line` citations.
- Brief note on next steps or likely primary file to inspect.

## Examples

- "Where is `PrefixJtiBlacklist` defined?" - search `config/` for the constant.
- "Which files reference `ACTIVATE_JWT`?" - grep across `config/`, `setTestEnv.sh`, `.env.sample`.

## Related Skills

- `code-navigation` (impact mapping), `file-reader` (detailed read)
