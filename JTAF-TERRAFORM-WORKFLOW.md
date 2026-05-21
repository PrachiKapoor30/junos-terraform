# JTAF Terraform Workflow - Setup Guide

## Current Status 

**COMPLETED:**
- ✅ NITA terminal connection established (66.129.234.205:40106)
- ✅ junos-terraform, yang, and venv cloned/setup on NITA
- ✅ JTAF tools verified (jtaf-pyang-plugindir, jtaf-provider, jtaf-yang2go)
- ✅ XML example configs available (dc1-spine1/2, dc1-leaf[1-3], dc1-borderleaf[1-2], dc1-firewall[1-2])
- ✅ Template files created (.terraformrc.template, etc-hosts.template)

**IN PROGRESS:**
- ⏳ Resolving YANG model recursion issue with pyang/18.2
- 🔧 Attempting lighter YANG versions (24.4, 25.4)

## YANG Model Processing Issue

The Junos 18.2 YANG models have deeply nested groupings causing Python recursion errors in pyang. This is a known limitation with pre-18.2 models.

**Solution: Using YANG 24.4 (newer, optimized)**

Run on NITA:
```bash
cd ~/junos-terraform
jtaf-yang2go -p ~/yang/24.4/24.4R1/common ~/yang/24.4/24.4R1/junos-qfx/conf/*.yang \
  -x examples/evpn-vxlan-dc/dc1/dc1-spine1.xml -t vqfx
```

## Manual Complete Workflow (if automated fails)

If YANG processing continues to fail, use manual JSON + provider approach:

```bash
# On NITA
cd ~/junos-terraform

# 1. Create minimal schema JSON
cat > minimal-schema.json << 'EOF'
{
  "root": {
    "name": "root",
    "type": "container",
    "children": [{"name": "configuration", "type": "container", "children": []}]
  },
  "identities": []
}
EOF

# 2. Generate provider with XML configs
jtaf-provider -j minimal-schema.json \
  -x examples/evpn-vxlan-dc/dc1/dc1-spine1.xml \
  -x examples/evpn-vxlan-dc/dc1/dc1-spine2.xml \
  -t vqfx

# 3. Build provider
cd terraform-provider-junos-vqfx && go install .
```

## JCL Device Connection Details (from CSV)

| Device | Alias | IP | SSH Port | User | Pass |
|--------|-------|----|----|------|------|
| dc1-spine1 | 003 dc1-spine1 | 66.129.234.205 | 40046 | jcluser | Juniper!1 |
| dc1-spine2 | 004 dc1-spine2 | 66.129.234.205 | 40054 | jcluser | Juniper!1 |
| dc1-leaf1 | 005 dc1-leaf1 | 66.129.234.205 | 40069 | jcluser | Juniper!1 |
| dc1-leaf2 | 006 dc1-leaf2 | 66.129.234.205 | 40080 | jcluser | Juniper!1 |
| dc1-leaf3 | 007 dc1-leaf3 | 66.129.234.205 | 40084 | jcluser | Juniper!1 |
| dc1-borderleaf1 | 001 dc1-borderleaf1 | 66.129.234.205 | 40003 | jcluser | Juniper!1 |
| dc1-borderleaf2 | 002 dc1-borderleaf2 | 66.129.234.205 | 40028 | jcluser | Juniper!1 |
| dc1-firewall1 | 001 dc1-firewall1 | 66.129.234.205 | 40007 | jcluser | Juniper!1 |
| dc1-firewall2 | 002 dc1-firewall2 | 66.129.234.205 | 40032 | jcluser | Juniper!1 |

## Next Steps

1. Try YANG 24.4 version for provider generation
2. If successful, build Go provider
3. Setup .terraformrc pointing to Go binary
4. Add device hostnames to /etc/hosts
5. Generate Terraform files with jtaf-xml2tf
6. Run terraform plan/apply

## Terraform Configuration Template

Place in `~/.terraformrc`:
```
provider_installation {
  dev_overrides {
    "registry.terraform.io/hashicorp/junos-vqfx" = "/home/jcluser/go/bin"
  }
  direct {}
}
```

## /etc/hosts Setup

Add to `/etc/hosts`:
```
66.129.234.205  dc1-spine1
66.129.234.205  dc1-spine2
66.129.234.205  dc1-leaf1
66.129.234.205  dc1-leaf2
66.129.234.205  dc1-leaf3
66.129.234.205  dc1-borderleaf1
66.129.234.205  dc1-borderleaf2
66.129.234.205  dc1-firewall1
66.129.234.205  dc1-firewall2
```

Then configure SSH ProxyCommand in ~/.ssh/config to route through correct ports.
