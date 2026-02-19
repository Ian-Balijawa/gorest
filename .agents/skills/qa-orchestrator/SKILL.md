---
name: qa-orchestrator
description: Orchestrate multi-step change workflows (search, analyze, propose, implement, verify) with explicit evidence and gate checks.
license: MIT
metadata:
  mode: workflow
  purpose: orchestrate
---

# QA Orchestrator

## When to Use

- Complex tasks that require multiple evidence-backed steps and human approval between stages.

## Workflow

1. **Locate:** `source-search` / `code-navigation`
2. **Understand:** `file-reader` / `explain-code`
3. **Simulate:** `safe-edit-simulator` (dry-run + risk checklist)
4. **Implement:** `patch-applier` + `code-formatter` (after approval)
5. **Verify:** `test-runner`, `linter-runner`, `static-analysis`

## Rules

- Keep each step small and reviewable; require explicit user approval for destructive or high-risk changes.
- Provide `path:line` evidence for key claims.

## Output

- Step-by-step results (numbered).
- Evidence for decisions with `path:line` citations.
- Next action and required approvals.

## Related Skills

- `safe-edit-simulator`, `patch-applier`, `test-runner`
