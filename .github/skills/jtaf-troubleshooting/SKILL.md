---
name: jtaf-troubleshooting
description: "Troubleshoot JTAF setup/build/provider issues including pyang plugin errors, Go/toolchain mismatch, Terraform plugin discovery, and build.sh or convert.sh failures."
argument-hint: "Error text or failing step (setup/build/convert/terraform)"
---

# JTAF Troubleshooting Skill

## When To Use
- Setup commands fail in Quick Start.
- `build.sh` or `convert.sh` fails.
- `pyang` or `jtaf-provider` commands fail.
- Terraform cannot find the custom Junos provider.

## Inputs
- Failing command.
- Error output.
- Current working directory.

## Procedure
1. Classify failure stage:
   - Environment/setup (Python, pip, venv)
   - Generation (`pyang`, plugindir, JSON output)
   - Provider build (`build.sh`, Go compile)
   - Convert/output (`convert.sh`, generated files)
   - Terraform runtime/provider discovery (`.terraformrc`, plugin path)
2. Reproduce minimally with one command.
3. Verify prerequisites:
   - Python and venv active
   - `pip install -e .` completed in `junos-terraform`
   - Go is available and compatible with provider module
   - Terraform installed when runtime checks are requested
4. Apply targeted fix from the signatures below.
5. Re-run failed step and report delta outcome.

## Error Signatures And Fixes

### `jtaf-pyang-plugindir: command not found`
- Ensure editable install completed in `junos-terraform` and venv is active.
- Re-run setup:
  - `python3 -m venv venv`
  - `. venv/bin/activate`
  - `pip install -e .`

### `pyang: command not found` or plugin format failure
- Install/repair pyang in active environment:
  - `pip install pyang`
- Re-run with explicit plugindir from `$(jtaf-pyang-plugindir)`.

### `go: command not found` or compile errors during provider build
- Confirm Go installation and PATH.
- From provider folder, run `go mod tidy` then retry build.
- Report exact compile error if symbol/type mismatch remains.

### `terraform init` cannot find custom provider
- Validate `~/.terraformrc` dev override path points to local `go/bin`.
- Confirm provider binary exists in configured path.
- Re-run `terraform init` after path correction.

### `build.sh`/`convert.sh` permission denied
- Add execute bit and retry:
  - `chmod +x build.sh convert.sh`

## Guardrails
- Keep fixes minimal and reversible.
- Do not run destructive git commands.
- Keep all commands aligned with repository workflow.

## Output Checklist
- Failure class and root cause.
- Exact fix command(s) run.
- Re-test result.
- Next fallback check if still failing.
