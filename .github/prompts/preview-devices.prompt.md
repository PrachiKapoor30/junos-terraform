---
description: Preview Terraform changes for one or more Junos devices by mode (preview vqfx, preview srx, preview all, or preview explicit targets) without applying.
---

Preview Terraform configuration changes by device scope.

Use when:
- You want a safe plan-only preview for VQFX, SRX, all devices, or explicit targets.

Accepted input forms (strict):
- /preview-devices preview vqfx
- /preview-devices preview qfx
- /preview-devices preview srx
- /preview-devices preview vsrx
- /preview-devices preview all
- /preview-devices preview targets <resource-address> [<resource-address> ...]

If input is not one of these forms, stop and ask user to choose one accepted form.

Execution rules:

1. Resolve working directory deterministically
- Required path: examples/terraform_files
- If current directory is not this path, change into it first.
- If path does not exist, stop and report missing path.

2. Resolve preview mode
- Treat vsrx as an alias of srx.
- Treat qfx as an alias of vqfx.
- For `preview vqfx`, `preview srx`, and `preview all`, use full `terraform plan` (no `-target`).
- Use `-target` only for `preview targets <resource...>` or explicit emergency troubleshooting.

Target set examples (only when using `preview targets` or emergency targeted mode):

VQFX targets:
- terraform-provider-junos-vqfx-evpn-vxlan.dc1-borderleaf1-base-config
- terraform-provider-junos-vqfx-evpn-vxlan.dc1-borderleaf2-base-config
- terraform-provider-junos-vqfx-evpn-vxlan.dc1-leaf1-base-config
- terraform-provider-junos-vqfx-evpn-vxlan.dc1-leaf2-base-config
- terraform-provider-junos-vqfx-evpn-vxlan.dc1-leaf3-base-config
- terraform-provider-junos-vqfx-evpn-vxlan.dc1-spine1-base-config
- terraform-provider-junos-vqfx-evpn-vxlan.dc1-spine2-base-config
- terraform-provider-junos-vqfx-evpn-vxlan.dc2-spine1-base-config
- terraform-provider-junos-vqfx-evpn-vxlan.dc2-spine2-base-config

SRX targets:
- terraform-provider-junos-vsrx-evpn-vxlan.dc1-firewall1-base-config
- terraform-provider-junos-vsrx-evpn-vxlan.dc1-firewall2-base-config
- terraform-provider-junos-vsrx-evpn-vxlan.dc2-firewall1-base-config
- terraform-provider-junos-vsrx-evpn-vxlan.dc2-firewall2-base-config

3. Run plan-only command
- Never run apply in this prompt.
- Do not run manual preflight connectivity checks (for example, nc).
- Let terraform plan be the reachability gate. If any target is unreachable, plan fails and you must stop.
- Show plan output directly in terminal/chat.
- Always save a binary plan file `preview.plan` and write full readable output to `preview_full_config.txt` in `examples/terraform_files`.
- Dependency checks (fail fast): verify `terraform` exists before running. If missing, stop and return `blocking_error`.
- Add run metadata header to `preview_full_config.txt` before plan text:
  - `generated_at` (UTC timestamp)
  - `mode`
  - `git_commit` (short SHA from current repo)

Mode: preview vqfx (or preview qfx)
terraform plan -no-color -out=preview.plan
{
  echo "# preview_metadata";
  echo "generated_at=$(date -u +%Y-%m-%dT%H:%M:%SZ)";
  echo "mode=preview-vqfx";
  echo "git_commit=$(git -C ../.. rev-parse --short HEAD 2>/dev/null || echo unknown)";
  echo;
  terraform show -no-color preview.plan;
} | tee preview_full_config.txt

Mode: preview srx (or preview vsrx)
terraform plan -no-color -out=preview.plan
{
  echo "# preview_metadata";
  echo "generated_at=$(date -u +%Y-%m-%dT%H:%M:%SZ)";
  echo "mode=preview-srx";
  echo "git_commit=$(git -C ../.. rev-parse --short HEAD 2>/dev/null || echo unknown)";
  echo;
  terraform show -no-color preview.plan;
} | tee preview_full_config.txt

Mode: preview all
terraform plan -no-color -out=preview.plan
{
  echo "# preview_metadata";
  echo "generated_at=$(date -u +%Y-%m-%dT%H:%M:%SZ)";
  echo "mode=preview-all";
  echo "git_commit=$(git -C ../.. rev-parse --short HEAD 2>/dev/null || echo unknown)";
  echo;
  terraform show -no-color preview.plan;
} | tee preview_full_config.txt

Mode: preview targets <resource...>
terraform plan -no-color -out=preview.plan -target='<resource-1>' -target='<resource-2>' ...
{
  echo "# preview_metadata";
  echo "generated_at=$(date -u +%Y-%m-%dT%H:%M:%SZ)";
  echo "mode=preview-targets";
  echo "git_commit=$(git -C ../.. rev-parse --short HEAD 2>/dev/null || echo unknown)";
  echo;
  terraform show -no-color preview.plan;
} | tee preview_full_config.txt

Post-run sanity checks:
- Verify `preview.plan` exists and is non-empty.
- Verify `preview_full_config.txt` exists and is non-empty.
- If either check fails, stop and return `blocking_error`.

Targeted mode warning:
- `-target` is not routine mode. Use only for explicit resource previews or recovery/troubleshooting.

4. Required output contract (compact by default)
- Return only 1 short summary line and 6 fields:
  mode, exit_code, plan_summary, terminal_output_shown, plan_file, full_output_file.
- Add warnings only if present.
- Add blocking_error only if present.
- Do not include extra command output unless user asks for full details.

Guardrails:
- Never run terraform apply in this prompt. If user requests apply, stop and redirect to /deploy-devices.
- Show plan output in terminal/chat directly.
- Create/update `preview.plan` and `preview_full_config.txt` only; do not create additional plan artifacts.
- Do not run terraform init when provider dev_overrides are active.
- Do not run extra commands or checks beyond the requested preview flow.
