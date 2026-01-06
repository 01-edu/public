#!/usr/bin/env bash
#Modified version for atos learners machines in the DevOps Cloud Branch

# The value of this parameter is expanded like PS1 and the expanded value is the
# prompt printed before the command line is echoed when the -x option is set
# (see The Set Builtin). The first character of the expanded value is replicated
# multiple times, as necessary, to indicate multiple levels of indirection.
# \D{%F %T} prints date like this : 2019-12-31 23:59:59
PS4='-\D{%F %T} '

# Print commands and their arguments as they are executed.
set -x

# Log stdout & stderr

# set bashrc path
echo "-------------------------------------------------------------------------------------"

echo "----------------------------------------------------------------------------------"

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
apt install gcc

# Install git
apt install git -y

# Install gcc
apt install gcc -y

# Install apt-get
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

# shellcheck disable=2016
sudo -i echo 'export PATH=$PATH:/usr/local/go/bin' >>/etc/profile
echo 'export PATH=\$PATH:/usr/local/go/bin:\$HOME/go/bin:$PATH' >>~/.bashrc

# Install Node.js

curl -sL https://deb.nodesource.com/setup_16.x | bash -
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

# Install LibreOffice

apt-get --no-install-recommends -y install libreoffice

# Install Chrome

wget 'https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb'
dpkg -i google-chrome-stable_current_amd64.deb
rm google-chrome-stable_current_amd64.deb

# End of script reached

# set bashrc path
echo "-------------------------------------------------------------------------------------"

echo "----------------------------------------------------------------------------------"

echo "basic-install-piscine script SUCCESSFUL"
