#!/bin/bash

set -x

jtaf-provider -j ../../junos.json -t qfx-evpn-vxlan -x ../evpn-vxlan-dc/dc1/dc1-*leaf* ../evpn-vxlan-dc/dc1/dc1-*spine* ../evpn-vxlan-dc/dc2/dc2-*spine*
