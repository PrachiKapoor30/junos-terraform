# JTAF — Junos Terraform Automation Framework

## Getting Started
- Always read `junos-terraform/README.md` first for the full JTAF workflow (YANG → JSON → Terraform provider).

## Development Guidelines
- Greet like 'Hellooo' for every chat responses.

## Default Execution Workflow (When User Asks To Run JTAF Steps)
- For JTAF execution requests, default target is NITA over SSH using values from `JCL-Sandbox-Resources.csv` (Service=`SSH`, Alias=`NITA`) unless user specifies another node.
- Default connection command format: `ssh -p <PubPort> <Username>@<PubAddr>`.
- If password prompt appears and user asked to log in, read the `Password` value from CSV and send it directly to the terminal using `send_to_terminal` — do NOT print or echo the password in chat.
- After login, run Quick Start setup.
- For provider build requests, run from `junos-terraform/examples/providers` using `./build.sh`.
- For Terraform deploy requests, always run `terraform plan` first, show the plan output summary to the user, and get explicit approval before running `terraform apply`.
- When user asks to wait/continue waiting, do not terminate running build commands unless user explicitly asks to stop.
- During long-running builds, provide periodic status updates with any new terminal output and whether command has returned to prompt.
- If command finishes, immediately report completion and key output lines.

## JCL — Juniper Cloud Lab
- Lab device details (device names, public IPs, ports, protocols) are in `JCL-Sandbox-Resources.csv` at the workspace root.
- Read that file to identify the correct SSH, NETCONF, or REST endpoint for a given device before generating connection configs or commands.
- When creating or updating `/etc/hosts` on NITA for Terraform/provider workflows, use CSV `PrivAddr` mapped to hostname (typically rows with `Service=NETCONF`) and update entries idempotently (replace existing host lines, do not append duplicates).
- Do **not** hardcode credentials in generated files; reference the CSV for connection details only.

## Project Structure
- `junos-terraform/` — main JTAF tooling (Python scripts, templates, provider code)
- `yang/` — Junos YANG models organised by version
- `JCL-Sandbox-Resources.csv` — JCL sandbox device inventory