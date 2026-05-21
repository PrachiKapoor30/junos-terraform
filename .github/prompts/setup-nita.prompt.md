---
name: Setup NITA
description: "Check if JTAF Quick Start setup is present on NITA and run it only for missing components."
argument-hint: "Optional: 'force' to re-run setup even if components exist"
agent: jtaf-agent
---

Check if the JTAF Quick Start setup is already present on NITA, then run only what is missing.

Requirements:
- Ensure an active SSH session to NITA exists. If not, run the Connect-JCL-SSH-Node prompt first.
- Before pre-flight, detect stale virtualenv state and repair it if needed:
  - Run:
    ```bash
    echo "$VIRTUAL_ENV"
    ls ~/venv 2>&1
    ```
  - If `$VIRTUAL_ENV` is set but `~/venv` is missing, reset shell state:
    ```bash
    deactivate 2>/dev/null || true
    unset VIRTUAL_ENV
    hash -r
    ```
- Run the following pre-flight check on NITA:
  ```bash
  ls ~/junos-terraform ~/yang ~/venv 2>&1
  ```
- Evaluate results:
  - `~/junos-terraform` missing → run `git clone https://github.com/juniper/junos-terraform`
  - `~/junos-terraform` exists → run `git -C ~/junos-terraform pull` to fetch latest changes before proceeding
  - `~/yang` missing → run `git clone https://github.com/juniper/yang`
  - `~/yang` exists → run `git -C ~/yang pull` to fetch latest changes
  - `~/venv` missing → run `python3 -m venv ~/venv`
- After any clones, activate venv and install dependencies exactly as per the README Quick Start:
  ```bash
  cd ~/junos-terraform
  . ~/venv/bin/activate
  pip install -e .
  ```
- If all three components are already present, report setup is complete and skip all install steps.
- If argument is `force`, re-run all setup steps regardless of existing state.

Report which components were found, which were installed, and final setup status.
