---
name: Deploy JTAF Provider
description: "Install built providers, configure .terraformrc and /etc/hosts on NITA, then run terraform plan/apply against JCL lab devices."
argument-hint: "Optional: 'plan-only' to skip apply, or a specific device hostname to target"
agent: jtaf-agent
---

Deploy the generated JTAF Terraform providers and apply configuration to JCL lab devices on NITA.

## Prerequisites check

Before deploying, verify on NITA:
1. Both provider directories exist under `~/junos-terraform/examples/providers/`:
   - `terraform-provider-junos-vqfx-evpn-vxlan/`
   - `terraform-provider-junos-vsrx-evpn-vxlan/`
2. `.tf` files exist in `~/junos-terraform/examples/terraform_files/`
3. If either is missing, instruct the user to run `/Run-JTAF-Provider-Build` first.

## Step 1 — Install providers

Run `go install .` in each provider directory:

```bash
cd ~/junos-terraform/examples/providers/terraform-provider-junos-vqfx-evpn-vxlan && go install .
cd ~/junos-terraform/examples/providers/terraform-provider-junos-vsrx-evpn-vxlan && go install .
```

Verify binaries exist:
```bash
ls ~/go/bin/terraform-provider-junos-*
```

## Step 2 — Configure .terraformrc

Copy the example `.terraformrc` to `~/.terraformrc` if not already present:
```bash
cp ~/junos-terraform/examples/example-terraformrc ~/.terraformrc
```

The file points both providers to `~/go/bin` — verify this matches the actual `go install` output path.

## Step 3 — Configure /etc/hosts

Add device hostname → IP entries to `/etc/hosts` on NITA so Terraform can resolve device names.
Read device IPs from `JCL-Sandbox-Resources.csv` using `PrivAddr` (Service=`NETCONF`).
Use `etc-hosts.template` in the workspace root as the reference mapping.

Important:
- Use private addresses (`PrivAddr`) for `/etc/hosts`, not public address (`PubAddr`).
- Apply updates idempotently: if host entries already exist, replace those host lines instead of appending duplicates.

```bash
sudo tee -a /etc/hosts << 'EOF'
<entries from etc-hosts.template>
EOF
```

Verify:
```bash
grep -E 'dc[12]-|wan-pe' /etc/hosts
```

## Step 4 — Set credentials and ports in .tf files

The generated `.tf` files in `~/junos-terraform/examples/terraform_files/` have empty `username`, `password`, and `port` fields.
- Read `Username` and `Password` from `JCL-Sandbox-Resources.csv` for each device.
- Read NETCONF `PubPort` for each device from the CSV (Service=`NETCONF`).
- Update the files using `sed` on NITA — do NOT print credentials in chat.

## Step 5 — Terraform plan

Always run with `-out` so the saved plan is used exactly by apply:

```bash
cd ~/junos-terraform/examples/terraform_files
terraform plan -out "tfplan"
```

Show the **full plan output** to the user (all resource diffs, not just the summary line).
Do not redirect to a log file — capture and display the complete output in chat.
Do not run `terraform apply` yet. Ask for explicit user approval first.

## Step 6 — Terraform apply

Unless `plan-only` argument was given:
```bash
terraform apply "tfplan"
```

Run this step only after explicit user approval following Step 5.
Poll silently until completion. Report final apply summary (resources applied, errors if any).

## Error handling

- `Plugin not found` → re-check `~/.terraformrc` path and `go install` output.
- `Connection refused` on NETCONF port → verify port numbers from CSV match `.tf` files.
- `Authentication failed` → verify credentials from CSV were applied correctly.
- Any other failure → consult `/jtaf-troubleshooting` skill.

## Reporting

Report concisely:
- Provider install: PASS/FAIL
- `.terraformrc`: present/updated
- `/etc/hosts`: entries added/already present
- `terraform plan`: summary line
- `terraform apply`: summary line or skipped
