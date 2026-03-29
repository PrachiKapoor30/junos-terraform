#!/bin/bash

set -e

jtaf-xml2yaml -x ../evpn-vxlan-dc/dc1/dc1-*leaf* ../evpn-vxlan-dc/dc1/dc1-*spine* ../evpn-vxlan-dc/dc2/dc2-*spine*  -j ansible-provider-junos-vqfx-ansible-role/trimmed_schema.json -d ansible_files --auto-detect-hierarchy --device-group-delta

jtaf-xml2yaml -x ../evpn-vxlan-dc/dc1/dc1-*firewall* ../evpn-vxlan-dc/dc2/dc2-*firewall*  -j ansible-provider-junos-srx-ansible-role/trimmed_schema.json -d ansible_files --auto-detect-hierarchy --device-group-delta

rm -rf ../../../ansible-evpn-vxlan-deploy/group_vars ../../../ansible-evpn-vxlan-deploy/host_vars && cp -R ansible_files/group_vars ansible_files/host_vars ../../ansible-evpn-vxlan-deploy/
