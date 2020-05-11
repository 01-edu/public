#!/usr/bin/env bash

# Install OpenSSH

script_dir="$(cd -P "$(dirname "$BASH_SOURCE")" && pwd)"
cd $script_dir
. set.sh

ssh_port=${1:-521}

# Install dependencies
apt-get -y install ssh

cat <<EOF>> /etc/ssh/sshd_config
Port $ssh_port
PasswordAuthentication no
AllowUsers root
EOF
