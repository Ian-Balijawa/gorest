---
name: explain-code
description: Produce concise, citation-backed explanations of module behavior, data flow, and important types/APIs.
license: MIT
metadata:
  mode: read-only
  purpose: explanation
---

# Explain Code

## When to Use

- The user asks "how does X work?" or needs a human-readable overview tied to exact code locations.

## Rules

- Read-only and source-cited: every non-trivial claim must include `path:line`.
- Keep summaries short and pragmatic. Focus on inputs/outputs, side effects, and error handling.

## Output

- **What it does:** 1-4 sentences.
- **How it works:** 3-6 bullets with `path:line` citations for key statements.
- **Key types/APIs:** short list with `path:line`.
- **Edge cases / invariants:** 1-3 bullets.

## Examples

- "Explain `handler.Login` behavior and where passwords are hashed."
- "How does the JWT blacklist work end-to-end?"

## Related Skills

- `file-reader` (low-level excerpts), `code-navigation` (impact mapping)
