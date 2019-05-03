#!/bin/bash

# Install firewall

SCRIPT_DIR="$(cd -P "$(dirname "$BASH_SOURCE")" && pwd)"
cd $SCRIPT_DIR
. set.sh

SSH_PORT=${1:-521}

apt-get -y install ufw

ufw logging off
ufw allow in "$SSH_PORT"/tcp
ufw --force enable
