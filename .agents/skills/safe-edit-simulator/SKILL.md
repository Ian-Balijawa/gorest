---
name: safe-edit-simulator
description: Produce a dry-run diff plus a risk checklist and rollback plan for any potentially disruptive change.
license: MIT
metadata:
  mode: planning
  purpose: safety
---

# Safe Edit Simulator

## When to Use

- Before applying changes that affect auth, security, DB migrations, middleware, or deployment behavior.

## Responsibilities

- Produce a readable dry-run diff and summarize changed hunks.
- Enumerate risks and mitigation steps.
- Provide a concise rollback plan and verification commands.

## Rules

- Do not apply changes; only simulate and document.
- Call out impacts on authentication, token validation, env/config, and data integrity.

## Output

- **Dry-run diff:** summary of files and hunks.
- **Risk checklist:** 3-8 items with severity.
- **Rollback:** single-line plan.
- **Verify:** 1-3 commands.

## Related Skills

- `fix-suggester`, `patch-applier`, `qa-orchestrator`
