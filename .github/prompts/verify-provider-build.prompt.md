---
name: Verify Provider Build
description: "Run post-build smoke checks for a generated Junos Terraform provider and return a pass/fail summary."
argument-hint: "Optional provider folder path and check scope"
agent: jtaf-agent
---

Run a standard post-build verification for the generated provider.

Requirements:
- Read `junos-terraform/README.md` first for expected build/install flow.
- Identify provider folder from argument or infer latest `terraform-provider-junos-*` directory.
- Run a concise smoke-check sequence and report pass/fail for each step.

Default smoke checks:
1. Confirm provider directory and expected key files exist (`go.mod`, `main.go`, provider files).
2. Run `go test ./...` in the provider folder.
3. Confirm provider binary presence or buildability (`go build ./...`).
4. Report generated/expected artifacts and any missing outputs.

Output format:
- `Provider path:`
- `Checks:`
  - `Structure:` PASS/FAIL
  - `Go tests:` PASS/FAIL
  - `Buildability:` PASS/FAIL
  - `Artifacts:` PASS/FAIL
- `Overall:` PASS/FAIL
- `Key failures:` brief actionable bullets

Argument handling:
- Parse argument as: `<provider_path> [quick|full]`.
- `quick` runs structure + buildability checks.
- `full` runs all checks.
