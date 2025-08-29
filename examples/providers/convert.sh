#!/bin/bash

set -x
set -e

cp terraform-provider-junos-qfx-evpn-vxlan/trimmed_schema.json .

QFXFILES=`echo ../evpn-vxlan-dc/dc1/dc1-*leaf* ../evpn-vxlan-dc/dc1/dc1-*spine* ../evpn-vxlan-dc/dc2/dc2-*spine*`

for i in $QFXFILES; do
	BASENAME=${i##*/}
	HN=${BASENAME%.xml}
	FN=${BASENAME%.xml}.tf
	jtaf-xml2tf -x $i -t qfx -n $HN > ../terraform_files/$FN
	sed -i '/Generated Terraform Configuration:/d' ../terraform_files/$FN
done

rm trimmed_schema.json

cp terraform-provider-junos-srx-evpn-vxlan/trimmed_schema.json .

SRXFILES=`echo ../evpn-vxlan-dc/dc1/dc1-*firewall* ../evpn-vxlan-dc/dc2/dc2-*firewall*`

for i in $SRXFILES; do
	BASENAME=${i##*/}
	HN=${BASENAME%.xml}
	FN=${BASENAME%.xml}.tf
	jtaf-xml2tf -x $i -t srx -n $HN > ../terraform_files/$FN
	sed -i '/Generated Terraform Configuration:/d' ../terraform_files/$FN
done

rm trimmed_schema.json
