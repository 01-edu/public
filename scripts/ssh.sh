#!/bin/bash

# Install OpenSSH

SCRIPT_DIR="$(cd -P "$(dirname "$BASH_SOURCE")" && pwd)"
cd $SCRIPT_DIR
. set.sh

SSH_PORT=${1:-521}

# Install dependencies
apt-get -y install ssh

cat <<EOF>> /etc/ssh/sshd_config
Port $SSH_PORT
PasswordAuthentication no
AllowUsers root
EOF
