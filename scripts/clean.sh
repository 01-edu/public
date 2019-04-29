#!/bin/bash

# Clean system

SCRIPT_DIR="$(cd -P "$(dirname "$BASH_SOURCE")" && pwd)"
cd $SCRIPT_DIR
. set.sh

# Purge useless packages
apt-get -y autoremove --purge
apt-get autoclean
apt-get clean
apt-get install

rm -rf /root/.local

# Remove connection logs
> /var/log/lastlog
> /var/log/wtmp
> /var/log/btmp

# Remove machine ID
> /etc/machine-id

# Remove logs
cd /var/log
rm -rf alternatives.log*
rm -rf apt/*
rm -rf auth.log
rm -rf dpkg.log*
rm -rf gpu-manager.log
rm -rf installer
rm -rf journal/d6e982aa8c9d4c1dbcbdcff195642300
rm -rf kern.log
rm -rf syslog
rm -rf sysstat

# Remove random seeds
rm -rf /var/lib/systemd/random-seed
rm -rf /var/lib/NetworkManager/secret_key

# Remove network configs
rm -rf /etc/NetworkManager/system-connections/*
rm -rf /var/lib/bluetooth/*
rm -rf /var/lib/NetworkManager/*

# Remove caches
rm -rf /var/lib/gdm3/.cache/*
rm -rf /root/.cache
rm -rf /home/student/.cache

rm -rf /home/student/.sudo_as_admin_successful /home/student/.bash_logout

rm -rf /tmp/*
rm -rf /tmp/.* || true
