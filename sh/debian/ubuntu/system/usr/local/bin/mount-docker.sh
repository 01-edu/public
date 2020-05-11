#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

cp -a /var/lib/docker /tmp
mount -t tmpfs -osize=2G tmpfs /var/lib/docker
mv /tmp/docker/* /var/lib/docker
