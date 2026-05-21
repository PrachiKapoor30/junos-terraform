---
name: JCL SSH Workflow
description: Use when connecting to JCL or Juniper Cloud Lab devices, reading JCL-Sandbox-Resources.csv, extracting public IP/port and credentials, or logging into NITA or other lab nodes over SSH.
applyTo: JCL-Sandbox-Resources.csv
---

# JCL SSH Workflow

- Read `JCL-Sandbox-Resources.csv` at the workspace root before generating any connection command.
- Find the requested device by alias or name, then use the row where `Service` is `SSH` unless the user explicitly asks for another protocol.
- Extract connection values from the CSV columns, not from the free-form `Url` text:
  - `PubAddr` for the public IP
  - `PubPort` for the public port
  - `Username` for the SSH username
  - `Password` for the SSH password when the user explicitly asks to log in
- Default SSH command format:

```bash
ssh -p <PubPort> <Username>@<PubAddr>
```

- Do not add `-o StrictHostKeyChecking=no` or `-o UserKnownHostsFile=/dev/null` unless the user explicitly asks for those flags.
- Run the SSH command in the terminal so the user can see the live terminal output side by side.
- If SSH prompts for a password and the user asked to log in, send the value from the CSV `Password` column.
- After sending the password, read terminal output and report whether login succeeded or failed.
- If the first connection attempt times out, check reachability with `nc -vz <PubAddr> <PubPort>` and report the result.
- When summarizing the result, include the exact CSV-derived connection details that were used.

## JTAF Execution Workflow

When a user asks to run JTAF steps or provider generation:
- **Default target** is NITA over SSH unless the user specifies another device
- Read `JCL-Sandbox-Resources.csv` to find NITA's connection details (Alias=`NITA`, Service=`SSH`)
- Use the connection command to initiate the SSH session
- After successful login, follow the setup and build steps from the terraform-provider instructions

## Example: NITA

- Alias: `NITA`
- Service: `SSH`
- Public IP: `66.129.234.205`
- Public port: `40106`
- Username: `jcluser`
- Password column value: `Juniper!1`