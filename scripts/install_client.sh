#!/bin/bash

# Configure Z01 client

# Log stdout & stderr
exec > >(tee -i /tmp/install_client.log) 2>&1

SCRIPT_DIR="$(cd -P "$(dirname "$BASH_SOURCE")" && pwd)"
cd $SCRIPT_DIR
. set.sh

# Set root password
passwd root

# Remove user password
passwd -d student
cp /etc/shadow /etc/shadow-

SSH_PORT=521
DISK=$(lsblk -o tran,kname,hotplug,type,fstype -pr |
	grep -e nvme -e sata |
	grep '0 disk' |
	cut -d' ' -f2 |
	sort |
	head -n1)

apt-get update
apt-get -y upgrade
apt-get -y autoremove --purge

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
sed -i -e "s|::DISK::|$DISK|g" etc/udev/rules.d/10-local.rules

apt-get -y install overlayroot
echo overlayroot=\"device:dev=/dev/disk/by-partlabel/01-tmp-system,recurse=0\" >> /etc/overlayroot.conf

# Fix permissions
find . -type d -exec chmod 755 {} \;
find . -type f -exec chmod 644 {} \;
find . -type f -exec /bin/sh -c "file {} | grep -q 'shell script' && chmod +x {}" \;
find . -type f -exec /bin/sh -c "file {} | grep -q 'public key' && chmod 400 {}" \;
cp --preserve=mode -RT . /

cd $SCRIPT_DIR
rm -rf /tmp/system

update-initramfs -u

. clean.sh
