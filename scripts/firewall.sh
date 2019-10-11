#!/usr/bin/env bash

# Install firewall

script_dir="$(cd -P "$(dirname "$BASH_SOURCE")" && pwd)"
cd $script_dir
. set.sh

SSH_PORT=${1:-521}

apt-get -y install ufw

ufw logging off
ufw allow in "$SSH_PORT"/tcp
ufw allow in 27960:27969
ufw --force enable
