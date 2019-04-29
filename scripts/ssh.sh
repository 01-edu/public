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

mkdir -p /root/.ssh
chmod -f 700 /root/.ssh
# echo 'ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIH30lZP4V26RVWWvAW91jM7UBSN68+xkuJc5cRionpMc' >> /root/.ssh/authorized_keys
chmod -f 600 /root/.ssh/authorized_keys || true

systemctl restart sshd.service
