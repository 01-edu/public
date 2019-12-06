#!/usr/bin/env bash

# Configure Z01 client

# Log stdout & stderr
exec > >(tee -i /tmp/install_client.log) 2>&1

script_dir="$(cd -P "$(dirname "$BASH_SOURCE")" && pwd)"
cd $script_dir
. set.sh

disk=$(lsblk -o tran,kname,hotplug,type,fstype -pr |
	grep -e nvme -e sata -e sas -e ata |
	grep '0 disk' |
	cut -d' ' -f2 |
	sort |
	head -n1)

systemctl stop unattended-upgrades.service

apt-get update
apt-get -y upgrade
apt-get -y autoremove --purge

# Remove outdated kernels
# old_kernels=$(ls -1 /boot/config-* | sed '$d' | xargs -n1 basename | cut -d- -f2,3)

# for old_kernel in $old_kernels; do
# 	dpkg -P $(dpkg-query -f '${binary:Package}\n' -W *"$old_kernel"*)
# done

apt-get -yf install

. bash_tweaks.sh
. ssh.sh
. firewall.sh
. ubuntu_tweaks.sh
. grub.sh "$disk"
. go.sh
. nodejs.sh
. fx.sh
. sublime.sh
. vscode.sh
. libreoffice.sh
. exam.sh

# Install additional packages
pkgs="
emacs
f2fs-tools
golang-mode
vim
xfsprogs
"

apt-get -y install $pkgs

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
