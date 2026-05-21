---
name: Run JTAF Provider Build
description: "Run the JTAF provider build workflow (Quick Start checks, build.sh, convert.sh) and summarize results."
argument-hint: "Optional device/provider target and whether to skip convert.sh"
agent: jtaf-agent
---

Run the JTAF provider build workflow for this workspace.

Requirements:
- Read `junos-terraform/README.md` first.
- Before building, verify setup is present using the Setup-NITA prompt. Run missing components only.
- Run build from `junos-terraform/examples/providers` using `./build.sh`.
- If build succeeds, run `./convert.sh` unless explicitly skipped.
- While `build.sh` or `convert.sh` are running, keep polling `get_terminal_output` silently until completion. Do NOT post intermediate "still running" chat messages — only report when a command finishes or errors.
- Report concise execution status with key output lines and next action.

If an argument is provided, treat it as constraints (for example: target provider, skip convert, or troubleshoot-only).
