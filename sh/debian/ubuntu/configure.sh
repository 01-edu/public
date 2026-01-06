#!/usr/bin/env bash

# Configure Z01 Ubuntu

set -euo pipefail
IFS='
'

# The value of this parameter is expanded like PS1 and the expanded value is the
# prompt printed before the command line is echoed when the -x option is set
# (see The Set Builtin). The first character of the expanded value is replicated
# multiple times, as necessary, to indicate multiple levels of indirection.
# \D{%F %T} prints date like this : 2019-12-31 23:59:59
PS4='-\D{%F %T} '

# Print commands and their arguments as they are executed.
set -x

# Log stdout & stderr
exec > >(tee -i /tmp/install_ubuntu.log) 2>&1

script_dir="$(cd -P "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$script_dir"

# Skip dialogs during apt-get install commands
export DEBIAN_FRONTEND=noninteractive # DEBIAN_PRIORITY=critical

export LC_ALL=C LANG=C
export SHELL=/bin/bash

disk=$(lsblk -o tran,kname,hotplug,type,fstype -pr |
	grep '0 disk' |
	cut -d' ' -f2 |
	sort |
	head -n1)

systemctl stop unattended-upgrades.service

apt-get --no-install-recommends update
apt-get --no-install-recommends -y upgrade
apt-get -y autoremove --purge

apt-get --no-install-recommends -y install curl

# Remove outdated kernels
# old_kernels=$(ls -1 /boot/config-* | sed '$d' | xargs -n1 basename | cut -d- -f2,3)

# for old_kernel in $old_kernels; do
# 	dpkg -P $(dpkg-query -f '${binary:Package}\n' -W *"$old_kernel"*)
# done

apt-get -yf install

# Configure Terminal

# Makes bash case-insensitive
cat <<EOF >>/etc/inputrc
set completion-ignore-case
set show-all-if-ambiguous On
set show-all-if-unmodified On
EOF

# Enhance Linux prompt
cat <<EOF >/etc/issue
Kernel build: \v
Kernel package: \r
Date: \d \t
IP address: \4
Terminal: \l@\n.\O

EOF

# Enable Bash completion
apt-get --no-install-recommends -y install bash-completion

cat <<EOF >>/etc/bash.bashrc
if ! shopt -oq posix; then
  if [ -f /usr/share/bash-completion/bash_completion ]; then
    . /usr/share/bash-completion/bash_completion
  elif [ -f /etc/bash_completion ]; then
    . /etc/bash_completion
  fi
fi
EOF

# Set-up all users
for dir in $(ls -1d /root /home/* 2>/dev/null || :); do
	# Hide login informations
	touch "$dir/.hushlogin"

	# Add convenient aliases & behaviors
	cat <<-'EOF' >>"$dir/.bashrc"
		export LS_OPTIONS="--color=auto"
		eval "`dircolors`"

		alias df="df --si"
		alias du="du --si"
		alias free="free -h --si"
		alias l="ls $LS_OPTIONS -al --si --group-directories-first"
		alias less="less -i"
		alias nano="nano -clDOST4"
		alias pstree="pstree -palU"

		HISTCONTROL=ignoreboth
		HISTFILESIZE=
		HISTSIZE=
		HISTTIMEFORMAT="%F %T "
	EOF

	# Fix rights
	usr=$(echo "$dir" | rev | cut -d/ -f1 | rev)
	chown -R "$usr:$usr" "$dir" || :
done

# Install OpenSSH

ssh_port=512

# Install dependencies
apt-get --no-install-recommends -y install ssh

cat <<EOF >>/etc/ssh/sshd_config
Port $ssh_port
PasswordAuthentication no
AllowUsers root
EOF

# Install firewall

apt-get --no-install-recommends -y install ufw

ufw logging off
ufw allow in "$ssh_port"/tcp
ufw allow in 27960:27969/tcp
ufw allow in 27960:27969/udp
ufw --force enable

# Install Grub

sed -i -e 's/message=/message_null=/g' /etc/grub.d/10_linux

cat <<EOF >>/etc/default/grub
GRUB_TIMEOUT=0
GRUB_RECORDFAIL_TIMEOUT=0
GRUB_TERMINAL=console
GRUB_DISTRIBUTOR=$()
GRUB_DISABLE_OS_PROBER=true
GRUB_DISABLE_SUBMENU=y
EOF

update-grub
grub-install "$disk"

# Install Go

wget https://dl.google.com/go/go1.16.3.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.16.3.linux-amd64.tar.gz
rm go1.16.3.linux-amd64.tar.gz
# shellcheck disable=2016
echo 'export PATH=$PATH:/usr/local/go/bin' >>/etc/profile

# Set-up all users
for dir in $(ls -1d /root /home/* 2>/dev/null || :); do
	# Add convenient aliases & behaviors
	cat <<-'EOF' >>"$dir/.bashrc"
		GOPATH=$HOME/go
		PATH=$PATH:$GOPATH/bin
		alias gobuild='CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags="-s -w"'
	EOF
	# shellcheck disable=2016
	echo 'GOPATH=$HOME/go' >>"$dir/.profile"

	# Fix rights
	usr=$(echo "$dir" | rev | cut -d/ -f1 | rev)
	chown -R "$usr:$usr" "$dir" || :
done

# Install Node.js

curl -sL https://deb.nodesource.com/setup_14.x | bash -
apt-get --no-install-recommends -y install nodejs

# Install FX: command-line JSON processing tool (https://github.com/antonmedv/fx)

npm install -g fx

# Install Sublime Text & Sublime Merge

wget -qO - https://download.sublimetext.com/sublimehq-pub.gpg | apt-key add -
apt-get --no-install-recommends install -y apt-transport-https

cat <<EOF >/etc/apt/sources.list.d/sublime-text.list
deb https://download.sublimetext.com/ apt/stable/
EOF

apt-get --no-install-recommends update
apt-get --no-install-recommends install -y sublime-text sublime-merge libgtk2.0-0

# Install Visual Studio Code

wget 'https://code.visualstudio.com/sha/download?build=stable&os=linux-deb-x64' --output-document vscode.deb
dpkg -i vscode.deb
rm vscode.deb

# Install VSCodium

wget -qO - https://gitlab.com/paulcarroty/vscodium-deb-rpm-repo/raw/master/pub.gpg | apt-key add -
echo 'deb https://paulcarroty.gitlab.io/vscodium-deb-rpm-repo/debs/ vscodium main' >>/etc/apt/sources.list.d/vscodium.list

apt-get --no-install-recommends update
apt-get --no-install-recommends install -y codium

# Set-up all users
for dir in $(ls -1d /home/* 2>/dev/null || :); do
	# Disable most of the telemetry and auto-updates
	mkdir -p "$dir/.config/Code/User"
	mkdir -p "$dir/.config/VSCodium/User"
	cat <<-'EOF' | tee \
		"$dir/.config/Code/User/settings.json" \
		"$dir/.config/VSCodium/User/settings.json"
			{
			    "gopls": {
			        "formatting.gofumpt": true
			    },
			    "extensions.autoCheckUpdates": false,
			    "extensions.autoUpdate": false,
			    "json.schemaDownload.enable": false,
			    "npm.fetchOnlinePackageInfo": false,
			    "settingsSync.keybindingsPerPlatform": false,
			    "telemetry.enableCrashReporter": false,
			    "telemetry.enableTelemetry": false,
			    "update.enableWindowsBackgroundUpdates": false,
			    "update.mode": "none",
			    "update.showReleaseNotes": false,
			    "workbench.enableExperiments": false,
			    "workbench.settings.enableNaturalLanguageSearch": false
			}
		EOF

	# Fix rights
	usr=$(echo "$dir" | rev | cut -d/ -f1 | rev)
	chown -R "$usr:$usr" "$dir" || :
done

# Install Go extension and tools

sudo -iu student code --install-extension golang.go
sudo -iu student go get github.com/01-edu/z01
sudo -iu student go get github.com/uudashr/gopkgs/v2/cmd/gopkgs
sudo -iu student go get github.com/ramya-rao-a/go-outline
sudo -iu student go get github.com/cweill/gotests/gotests
sudo -iu student go get github.com/fatih/gomodifytags
sudo -iu student go get github.com/josharian/impl
sudo -iu student go get github.com/haya14busa/goplay/cmd/goplay
sudo -iu student go get github.com/go-delve/delve/cmd/dlv
sudo -iu student go get github.com/go-delve/delve/cmd/dlv@master
sudo -iu student go get honnef.co/go/tools/cmd/staticcheck
sudo -iu student go get golang.org/x/tools/gopls
sudo -iu student go get mvdan.cc/gofumpt

# Install LibreOffice

apt-get --no-install-recommends -y install libreoffice

# Install Docker

apt-get --no-install-recommends -y install apt-transport-https ca-certificates curl gnupg2 software-properties-common
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add -
add-apt-repository --yes "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
apt-get --no-install-recommends update
apt-get --no-install-recommends -y install docker-ce docker-ce-cli containerd.io
adduser student docker

# Install Docker compose
curl -L "https://github.com/docker/compose/releases/download/1.29.1/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose
curl -L https://raw.githubusercontent.com/docker/compose/1.29.1/contrib/completion/bash/docker-compose -o /etc/bash_completion.d/docker-compose

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

# shellcheck disable=2086
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
i3lock
imagemagick
mpv
vim
virtualbox
xfsprogs
zenity
"
# shellcheck disable=2086
apt-get --no-install-recommends -y install $pkgs

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
# shellcheck disable=2086
systemctl disable $services

services="
grub-common.service
plymouth-quit-wait.service
"
# shellcheck disable=2086
systemctl mask $services

# Logout quickly
cat <<EOF >>/etc/systemd/logind.conf
KillUserProcesses=yes
UserStopDelaySec=0
EOF

# Disable GTK hidden scroll bars
echo GTK_OVERLAY_SCROLLING=0 >>/etc/environment

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
swapoff /swapfile || :
rm -f /swapfile
sed -i '/swapfile/d' /etc/fstab

# Put temporary and cache folders as tmpfs
echo 'tmpfs /tmp tmpfs defaults,noatime,rw,nosuid,nodev,mode=1777,size=1G 0 0' >>/etc/fstab

# Install additional drivers
ubuntu-drivers install || :

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

cd /usr/local/src/format
PATH=$PATH:/usr/local/go/bin
go mod download
go build -o /usr/local/bin/format

cd "$script_dir"
rm -rf /tmp/system

if ! test -v PERSISTENT; then
	sgdisk --new 0:0:+32G "$disk"
	sgdisk --new 0:0:+32G "$disk"
	sgdisk --largest-new 0 "$disk"
	sgdisk --change-name 3:01-tmp-home "$disk"
	sgdisk --change-name 4:01-docker "$disk"
	sgdisk --change-name 5:01-tmp-system "$disk"

	# Add Docker persistent partition
	partprobe
	mkfs.ext4 -E lazy_journal_init,lazy_itable_init=0 /dev/disk/by-partlabel/01-docker
	echo 'PARTLABEL=01-docker /var/lib/docker ext4 noatime,errors=remount-ro 0 2' >>/etc/fstab
	systemctl stop docker.service containerd.service
	mv /var/lib/docker /tmp
	mkdir /var/lib/docker
	mount /dev/disk/by-partlabel/01-docker
	mv /tmp/docker/* /var/lib/docker
	umount /var/lib/docker

	# Remove fsck because the system partition will be read-only (overlayroot)
	rm /usr/share/initramfs-tools/hooks/fsck

	apt-get --no-install-recommends -y install overlayroot
	echo 'overlayroot="device:dev=/dev/disk/by-partlabel/01-tmp-system,recurse=0"' >>/etc/overlayroot.conf

	update-initramfs -u

	# Lock root password
	passwd -l root

	# Disable user password
	passwd -d student

	# Remove tty
	cat <<-"EOF" >>/etc/systemd/logind.conf
		NAutoVTs=0
		ReserveVT=N
	EOF

	# Remove user abilities
	sed -i 's/^%admin/# &/' /etc/sudoers
	sed -i 's/^%sudo/# &/' /etc/sudoers
	gpasswd -d student lpadmin
	gpasswd -d student sambashare

	# Give to rights to use format tool
	echo 'student ALL = (root) NOPASSWD: /usr/local/bin/format' >>/etc/sudoers

	cp /etc/shadow /etc/shadow-
fi

# Use Cloudflare DNS server
echo 'supersede domain-name-servers 1.1.1.1;' >>/etc/dhcp/dhclient.conf

# Clean system

# Purge useless packages
apt-get -y autoremove --purge
apt-get autoclean
apt-get clean
apt-get install

rm -rf /root/.local

# Remove connection logs
echo >/var/log/lastlog
echo >/var/log/wtmp
echo >/var/log/btmp

# Remove machine ID
echo >/etc/machine-id

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
rm -rf /tmp/.* || :
