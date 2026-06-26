---
description: Connect to JCL/NITA over SSH using JCL-Sandbox-Resources.csv resolution.
---

Connect to a JCL device (default alias NITA) using inventory-driven values.

Use when:
- You want the agent to open an SSH session to JCL/NITA with resolved host, port, and username.

Accepted input forms (strict):
- /connect-jcl-nita connect nita
- /connect-jcl-nita connect alias <alias>

If input is not one of these forms, stop and ask user to choose one accepted form.

Execution rules:

1. Resolve inventory and target row
- Required inventory file: junos-terraform/JCL-Sandbox-Resources.csv
- If file is missing, stop and return blocking_error.
- Resolve alias:
  - `connect nita` -> alias is `NITA`
  - `connect alias <alias>` -> alias is `<alias>`
- From CSV, select row where:
  - `Alias` equals resolved alias (case-insensitive), and
  - `Service` equals `SSH`
- If no matching SSH row exists, stop and return blocking_error.

2. Resolve connection values from CSV columns
- `host` = `PubAddr`
- `port` = `PubPort`
- `username` = `Username`
- `password` = `Password`
- Build SSH command exactly as:
  - `ssh -p <port> <username>@<host>`

3. Connect
- Run only the resolved SSH command.
- If prompted for password, send the resolved `password` value.
- Do not run extra diagnostics or preflight checks unless user explicitly asks.

4. Required output contract (compact)
- Return 1 short summary line and fields:
  - alias
  - service
  - host
  - port
  - username
  - connected
  - exit_code
  - terminal_output_shown
- Add warnings only if present.
- Add blocking_error only if present.

Guardrails:
- Do not hardcode endpoint values when CSV lookup is available.
- Do not print password in response fields.
- Do not run unrelated commands before or after SSH connection.