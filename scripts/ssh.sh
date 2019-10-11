#!/usr/bin/env bash

# Install OpenSSH

script_dir="$(cd -P "$(dirname "$BASH_SOURCE")" && pwd)"
cd $script_dir
. set.sh

SSH_PORT=${1:-521}

# Install dependencies
apt-get -y install ssh

cat <<EOF>> /etc/ssh/sshd_config
Port $SSH_PORT
PasswordAuthentication no
AllowUsers root
EOF
