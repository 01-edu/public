#!/usr/bin/env bash

# Configure ubuntu desktop systems

script_dir="$(cd -P "$(dirname "$BASH_SOURCE")" && pwd)"
cd "$script_dir"
. set.sh

# Install dependencies
apt-get -y install lz4

# Change ext4 default mount options
sed -i -e 's/ errors=remount-ro/ noatime,nodelalloc,errors=remount-ro/g' /etc/fstab

# Disable GTK hidden scroll bars
echo GTK_OVERLAY_SCROLLING=0 >> /etc/environment

# Reveal boot messages
sed -i -e 's/TTYVTDisallocate=yes/TTYVTDisallocate=no/g' /etc/systemd/system/getty.target.wants/getty@tty1.service

# Speedup boot
sed -i 's/MODULES=most/MODULES=dep/g' /etc/initramfs-tools/initramfs.conf
sed -i 's/COMPRESS=gzip/COMPRESS=lz4/g' /etc/initramfs-tools/initramfs.conf

# Reveal autostart services
sed -i 's/NoDisplay=true/NoDisplay=false/g' /etc/xdg/autostart/*.desktop

# Remove password complexity constraints
sed -i 's/ obscure / minlen=1 /g' /etc/pam.d/common-password

# Remove splash screen (plymouth)
sed -i 's/quiet splash/quiet/g' /etc/default/grub

update-initramfs -u
update-grub

# Disable swapfile
swapoff /swapfile ||:
rm -f /swapfile
sed -i '/swapfile/d' /etc/fstab


# Purge unused Ubuntu packages
pkgs="
apport
bind9
bolt
cups*
exim*
fprintd
friendly-recovery
gnome-initial-setup
gnome-online-accounts
gnome-power-manager
gnome-software
gnome-software-common
memtest86+
orca
popularity-contest
python3-update-manager
secureboot-db
snapd
speech-dispatcher*
spice-vdagent
ubuntu-report
ubuntu-software
unattended-upgrades
update-inetd
update-manager-core
update-notifier
update-notifier-common
whoopsie
xdg-desktop-portal
"

apt-get -y purge $pkgs
apt-get -y autoremove --purge

services="
apt-daily-upgrade.timer
apt-daily.timer
console-setup.service
e2scrub_reap.service
keyboard-setup.service
motd-news.timer
remote-fs.target
"
systemctl disable $services

services="
grub-common.service
plymouth-quit-wait.service
"
systemctl mask $services

# Install packages
pkgs="$(cat common_packages.txt)
baobab
blender
dconf-editor
firefox
gimp
gnome-calculator
gnome-system-monitor
gnome-tweaks
i3lock
imagemagick
mpv
zenity
"

# Replace debian packages with ubuntu's
pkgs=${pkgs/linux-image-amd64/linux-image-generic}
pkgs=${pkgs/linux-headers-amd64/linux-headers-generic}
pkgs=${pkgs/firmware-linux-nonfree}

apt-get -y install $pkgs
