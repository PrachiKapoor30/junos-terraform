---
name: Connect JCL SSH Node
description: "Resolve SSH connection details from JCL-Sandbox-Resources.csv and connect to a JCL node (default NITA)."
argument-hint: "Device alias/name and login intent"
agent: jtaf-agent
---

Connect to a JCL node over SSH using the lab inventory CSV.

Requirements:
- Read `JCL-Sandbox-Resources.csv` and find the row by alias/name.
- Default to `Alias=NITA` and `Service=SSH` if no target is given.
- Use CSV columns `PubAddr`, `PubPort`, and `Username` for command construction.
- Run `ssh -p <PubPort> <Username>@<PubAddr>`.
- If asked to log in and password is prompted, read the `Password` from CSV and send it directly to the terminal via `send_to_terminal` — never print or display the password in chat.
- If SSH fails, check `nc -vz <PubAddr> <PubPort>` and report result.

If an argument is provided, use it to select target and behavior (connect-only vs full login).
