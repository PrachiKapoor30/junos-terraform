---
name: jtaf-provider-build
description: "Run JTAF provider setup/build/convert workflow from junos-terraform examples with README-aligned steps and status reporting. Use for build.sh, convert.sh, provider artifacts, and troubleshooting build output."
argument-hint: "Provider build request, device type, and whether convert.sh should run"
---

# JTAF Provider Build Skill

## When To Use
- User asks to build or rebuild JTAF Terraform providers.
- User asks to run `junos-terraform/examples/providers/build.sh`.
- User asks to run `convert.sh` after provider generation.
- Build output must be monitored and summarized.

## Required Inputs
- Optional target device/provider flavor if user specifies one.
- Optional instruction to skip convert step.

## Procedure
1. Read `junos-terraform/README.md` first.
2. Run only Quick Start setup commands from README when environment is not prepared.
3. Change into `junos-terraform/examples/providers`.
4. Run `./build.sh` and monitor for completion.
5. If build succeeds, run `./convert.sh` unless user asked to skip it.
6. Summarize key output: success/failure, generated artifacts, and next actions.

## Guardrails
- Do not run extra generation commands unless user explicitly asks.
- Do not terminate long-running build commands when user asks to keep waiting.
- Report whether terminal returned to prompt.
- Keep outputs concise and actionable.

## Output Checklist
- Commands executed.
- Build result and relevant key lines.
- Convert result and relevant key lines.
- Suggested next step when failures occur.
