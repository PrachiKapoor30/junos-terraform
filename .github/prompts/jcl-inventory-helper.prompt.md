---
name: JCL Inventory Helper
description: "Query JCL-Sandbox-Resources.csv by alias/service and return ready-to-run connection commands for SSH/NETCONF/REST endpoints."
argument-hint: "Alias/device, service, and optional connect-now flag"
agent: jtaf-agent
---

Resolve lab endpoint details from `JCL-Sandbox-Resources.csv`.

Requirements:
- Read CSV at workspace root and match rows by alias or device name.
- If service not provided, prefer `SSH`.
- Return structured output with: Alias, Service, PubAddr, PubPort, Username.
- Provide ready-to-run command snippets:
  - SSH: `ssh -p <PubPort> <Username>@<PubAddr>`
  - NETCONF (example): `ssh -p <PubPort> <Username>@<PubAddr> -s netconf`
  - REST (example URL): `https://<PubAddr>:<PubPort>/`
- If argument includes `connect-now`, run the SSH command for matched SSH endpoint.
- Never hardcode or invent connection values.

Argument handling:
- Parse argument as: `<alias_or_name> [service] [connect-now]`.
