---
name: jcl-ssh-connect
description: "Connect to JCL devices over SSH using JCL-Sandbox-Resources.csv as source of truth. Use for NITA/default login, extracting PubAddr/PubPort/Username, and authenticated SSH session checks."
argument-hint: "Device alias/name and whether to authenticate"
---

# JCL SSH Connect Skill

## When To Use
- User asks to connect or log in to JCL/Juniper Cloud Lab nodes.
- User asks for SSH command from lab inventory.
- User asks to access NITA by default.

## Required Inputs
- Device alias/name. Default is `NITA` when not specified.
- Protocol intent. Default is `SSH`.
- Whether user wants full login (password submission).

## Procedure
1. Read `JCL-Sandbox-Resources.csv` at workspace root.
2. Find row by device alias/name and `Service=SSH` unless user requests another protocol.
3. Extract `PubAddr`, `PubPort`, `Username`, and `Password` from CSV columns.
4. Build command: `ssh -p <PubPort> <Username>@<PubAddr>`.
5. Run command in terminal and wait for prompt.
6. If password prompt appears and user requested login, read the `Password` value from CSV and send it directly to the terminal using `send_to_terminal` — do NOT print or display the password in chat.
7. If connection fails or times out, run `nc -vz <PubAddr> <PubPort>` and report result.

## Guardrails
- Do not hardcode credentials.
- Do not use free-form `Url` field for command construction.
- Do not add insecure SSH flags unless user explicitly requests them.
- Never print, echo, or display the CSV `Password` value in chat — always send it silently and directly via `send_to_terminal`.

## Output Checklist
- Exact CSV-derived connection fields used.
- Login success/failure status.
- Reachability test result if SSH fails.
