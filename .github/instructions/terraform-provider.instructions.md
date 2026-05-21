name: Terraform Provider Build
description: Use when building Terraform providers for Junos; read README, run setup, and execute build.sh and convert.sh.
applyTo:
  - junos-terraform/terraform_provider/**/*
  - junos-terraform/examples/providers/**/*
  - terraform-provider-junos-*/**/*
  - "*.tf"
---

# Terraform Provider Build Workflow

Follow only these steps.

<!-- ## 1) Read the README first
```bash
cd junos-terraform
``` -->

## 1) Run Quick start & setup from Readme
<!-- ```bash
cd junos-terraform
. ../venv/bin/activate
pip install -e .
``` -->
- Run the commands mentioned in the "junos-terraform/README.md"  - "Quick Start" section to set up the Junos-Terraform Environment and Workflow
- Don't run any other commands from the README unless the user explicitly asks for them. Focus on the Quick Start and setup steps first before provider generation.

## 2) Build providers
```bash
cd junos-terraform/examples/providers
./build.sh
```

## 3) Convert generated configs
```bash
./convert.sh
```

Do not use extra provider generation steps unless the user explicitly asks for them.
Complete build first and then report the result before moving on to any other steps.
Run convert.sh only after build.sh completes successfully and report the result.
