---
name: Generate JTAF Provider E2E
description: "Generate Terraform provider artifacts end-to-end from YANG and XML inputs (YANG -> JSON -> provider) using JTAF workflow."
argument-hint: "Junos version, yang/common paths, xml glob, and device type"
agent: jtaf-agent
---

Run an end-to-end JTAF provider generation workflow.

Requirements:
- Read `junos-terraform/README.md` first.
- Follow the documented sequence: YANG -> JSON -> Terraform provider.
- Use `pyang` with `$(jtaf-pyang-plugindir)` to generate JSON.
- Use `jtaf-provider` with generated JSON, XML input(s), and device type.
- If arguments are missing, infer from repository examples and clearly state assumptions.
- Summarize exact commands used, output location, and key success/failure lines.

Argument handling:
- Parse argument as: `<junos_version> <common_path> <yang_glob> <xml_glob> <device_type>`.
- If argument includes `--dry-run`, only print commands that would run.
