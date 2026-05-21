# JTAF Copilot Customization Spec

## Goal
Provide reusable Copilot assets for the most common JTAF operator tasks:
- Building Terraform providers
- Connecting to JCL nodes over SSH

## Included Assets
- Skills:
  - `.github/skills/jtaf-provider-build/SKILL.md`
  - `.github/skills/jcl-ssh-connect/SKILL.md`
- Prompts:
  - `.github/prompts/run-jtaf-provider-build.prompt.md`
  - `.github/prompts/connect-jcl-ssh.prompt.md`

## Behavioral Requirements
- Always read `junos-terraform/README.md` before JTAF workflow execution.
- Use `JCL-Sandbox-Resources.csv` as source of truth for JCL connectivity values.
- Default JCL target is NITA over SSH unless user overrides it.
- Provider build path is `junos-terraform/examples/providers` with `./build.sh` first.

## Acceptance Criteria
- Skills and prompts are discoverable via `/` in chat.
- Prompt descriptions clearly indicate usage intent.
- Skill names match their folder names.
- Added assets are workspace-scoped under `.github/`.

## Future Extensions
- Add a prompt for full YANG -> JSON -> provider generation with user-selected Junos version.
- Add a troubleshooting skill for common build and plugin installation failures.
