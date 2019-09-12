#!/bin/bash

# Configure Z01 client

# Log stdout & stderr
exec > >(tee -i /tmp/install_client.log) 2>&1

SCRIPT_DIR="$(cd -P "$(dirname "$BASH_SOURCE")" && pwd)"
cd $SCRIPT_DIR
. set.sh

SSH_PORT=521
DISK=$(lsblk -o tran,kname,hotplug,type,fstype -pr |
	grep -e nvme -e sata |
	grep '0 disk' |
	cut -d' ' -f2 |
	sort |
	head -n1)

systemctl stop unattended-upgrades.service

sgdisk -n0:0:+32G "$DISK"
sgdisk -N0 "$DISK"
sgdisk -c3:01-tmp-home "$DISK"
sgdisk -c4:01-tmp-system "$DISK"

apt-get update
apt-get -y upgrade
apt-get -y autoremove --purge

# Remove outdated kernels
old_kernels=$(ls -1 /boot/config-* | xargs -n1 basename | grep -v "$(uname -r)" | cut -d- -f2,3)

for old_kernel in $old_kernels; do
	dpkg -P $(dpkg-query -f '${binary:Package}\n' -W *"$old_kernel"*)
done

. bash_tweaks.sh
. ssh.sh
. firewall.sh
. ubuntu_tweaks.sh
. grub.sh "$DISK"
. go.sh
. nodejs.sh
. fx.sh
. sublime.sh
. vscode.sh
. libreoffice.sh

# Install additional packages
PKGS="
emacs
f2fs-tools
golang-mode
vim
xfsprogs
"

apt-get -y install $PKGS

# Install additional drivers
ubuntu-drivers install

# Remove fsck because the system partition will be read-only (overlayroot)
rm /usr/share/initramfs-tools/hooks/fsck

# Copy system files

cp -r system /tmp
cd /tmp/system

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

sed -i -e "s|::DISK::|$DISK|g" etc/udev/rules.d/10-local.rules

# Generate wallpaper
cd usr/share/backgrounds/01
test ! -e wallpaper.png && composite logo.png background.png wallpaper.png
cd /tmp/system

cp --preserve=mode -RT . /

cd $SCRIPT_DIR
rm -rf /tmp/system

apt-get -y install overlayroot
echo overlayroot=\"device:dev=/dev/disk/by-partlabel/01-tmp-system,recurse=0\" >> /etc/overlayroot.conf

update-initramfs -u

# Remove root & user password
passwd -d root
passwd -d student
cp /etc/shadow /etc/shadow-

# Remove user ability to use sudo
gpasswd -d student sudo

. clean.sh
