#!/usr/bin/env bash

# Configure Z01 Ubuntu

# Log stdout & stderr
exec > >(tee -i /tmp/install_ubuntu.log) 2>&1

script_dir="$(cd -P "$(dirname "$BASH_SOURCE")" && pwd)"
cd $script_dir
. set.sh

disk=$(lsblk -o tran,kname,hotplug,type,fstype -pr |
	grep '0 disk' |
	cut -d' ' -f2 |
	sort |
	head -n1)

systemctl stop unattended-upgrades.service

apt-get update
apt-get -y upgrade
apt-get -y autoremove --purge

apt-get -y install curl

# Remove outdated kernels
# old_kernels=$(ls -1 /boot/config-* | sed '$d' | xargs -n1 basename | cut -d- -f2,3)

# for old_kernel in $old_kernels; do
# 	dpkg -P $(dpkg-query -f '${binary:Package}\n' -W *"$old_kernel"*)
# done

apt-get -yf install

. bash_tweaks.sh
. ssh.sh
. firewall.sh
. grub.sh "$disk"
. go.sh
. nodejs.sh
. fx.sh
. sublime.sh
. vscode.sh
. libreoffice.sh
. exam.sh
. docker.sh

# Purge unused Ubuntu packages
pkgs="
apparmor
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

# Install packages
pkgs="$(cat common_packages.txt)
baobab
blender
dconf-editor
emacs
f2fs-tools
firefox
gimp
gnome-calculator
gnome-system-monitor
gnome-tweaks
golang-mode
i3lock
imagemagick
mpv
vim
virtualbox
xfsprogs
zenity
"
apt-get -y install $pkgs

# Disable services
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

# Change ext4 default mount options
sed -i -e 's/ errors=remount-ro/ noatime,nodelalloc,errors=remount-ro/g' /etc/fstab

# Disable swapfile
swapoff /swapfile ||:
rm -f /swapfile
sed -i '/swapfile/d' /etc/fstab

# Put temporary and cache folders as tmpfs
echo 'tmpfs /tmp tmpfs defaults,noatime,rw,nosuid,nodev,noexec,mode=1777,size=1G 0 0' >> /etc/fstab

# Install additional drivers
ubuntu-drivers install ||:

# Copy system files

cp -r system /tmp
cd /tmp/system

test -v PERSISTENT && rm -rf etc/gdm3 usr/share/initramfs-tools

# Overwrite with custom files from Git repository
if test -v OVERWRITE; then
	folder=$(echo "$OVERWRITE" | cut -d';' -f1)
	url=$(echo "$OVERWRITE" | cut -d';' -f2)
	if git ls-remote -q "$url" &>/dev/null; then
		tmp=$(mktemp -d)
		git clone --depth 1 "$url" "$tmp"
		rm -rf "$tmp"/.git
		cp -aT "$tmp" "$folder"
		rm -rf "$tmp"
	fi
fi

# Fix permissions
find . -type d -exec chmod 755 {} \;
find . -type f -exec chmod 644 {} \;
find . -type f -exec /bin/sh -c "file {} | grep -q 'shell script' && chmod +x {}" \;
find . -type f -exec /bin/sh -c "file {} | grep -q 'public key' && chmod 400 {}" \;

sed -i -e "s|::DISK::|$disk|g" etc/udev/rules.d/10-local.rules

# Generate wallpaper
cd usr/share/backgrounds/01
test ! -e wallpaper.png && composite logo.png background.png wallpaper.png
cd /tmp/system

cp --preserve=mode -RT . /

cd $script_dir
rm -rf /tmp/system

if ! test -v PERSISTENT; then
	sgdisk -n0:0:+32G "$disk"
	sgdisk -N0 "$disk"
	sgdisk -c3:01-tmp-home "$disk"
	sgdisk -c4:01-tmp-system "$disk"

	# Remove fsck because the system partition will be read-only (overlayroot)
	rm /usr/share/initramfs-tools/hooks/fsck

	apt-get -y install overlayroot
	echo 'overlayroot="device:dev=/dev/disk/by-partlabel/01-tmp-system,recurse=0"' >> /etc/overlayroot.conf

	update-initramfs -u

	# Lock root password
	passwd -l root

	# Disable user password
	passwd -d student

	# Enable docker relocation
	systemctl enable mount-docker

	# Remove tty
	cat <<-EOF>> /etc/systemd/logind.conf
	NAutoVTs=0
	ReserveVT=N
	EOF

	# Remove user abilities
	gpasswd -d student sudo
	gpasswd -d student lpadmin
	gpasswd -d student sambashare

	cp /etc/shadow /etc/shadow-
fi

. clean.sh
