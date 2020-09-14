#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

# Fix Debian 10 bug (https://bugs.debian.org/cgi-bin/bugreport.cgi?bug=905409)
PATH=/sbin:/usr/sbin:$PATH

# Debian stable OS
apt-get update
apt-get -y upgrade
apt-get -y dist-upgrade

# Disable OpenStack SSH malware
mv /home/debian/.ssh/authorized_keys /root/.ssh/authorized_keys ||:
sed -i '/Generated-by-Nova/d' /root/.ssh/authorized_keys ||:
chown root:root /root/.ssh/authorized_keys ||:

# Terminal goodies
touch .hushlogin

cat <<'EOF'>> /root/.bashrc
export LS_OPTIONS="--color=auto"
eval "`dircolors`"

alias ctop="docker run --rm -it --name=ctop -v /var/run/docker.sock:/var/run/docker.sock:ro quay.io/vektorlab/ctop"
alias df="df --si"
alias du="du -cs --si"
alias free="free -h --si"
alias l="ls $LS_OPTIONS -al --si --group-directories-first"
alias less="less -i"
alias nano="nano -clDOST4"
alias pstree="pstree -palU"
alias gobuild='CGO_ENABLED=0 GOARCH=amd64 go build -trimpath -ldflags="-s -w"'

export HISTFILESIZE=
export HISTSIZE=
export HISTTIMEFORMAT="%F %T "

GOPATH=$HOME/go
HISTCONTROL=ignoreboth
HISTFILESIZE=
HISTSIZE=
HISTTIMEFORMAT="%F %T "
EOF

cat <<'EOF'>> /etc/inputrc
set completion-ignore-case
set show-all-if-ambiguous On
set show-all-if-unmodified On
EOF

cat <<'EOF'>> /etc/bash.bashrc
if ! shopt -oq posix; then
  if [ -f /usr/share/bash-completion/bash_completion ]; then
    . /usr/share/bash-completion/bash_completion
  elif [ -f /etc/bash_completion ]; then
    . /etc/bash_completion
  fi
fi
EOF

# Basic packages
apt-get -y install man bash-completion git ufw jq curl build-essential netcat wget psmisc lz4 file net-tools brotli unzip zip moreutils xauth sysfsutils rsync iperf pv tree mc screen ssh iotop whois sudo

# Enable time synchronization
timedatectl set-ntp true

# Configure screen
cat <<'EOF'>> /etc/screenrc
startup_message off
shell -$SHELL
defscrollback 100000
bind l eval clear "scrollback 0" "scrollback 100000"
EOF

# Configure SSH
cat <<'EOF'>> /etc/ssh/sshd_config
Port 521
PasswordAuthentication no
AllowUsers root
X11UseLocalhost no
EOF
systemctl restart ssh

touch /root/.Xauthority

# Firewall
ufw allow in 80/tcp
ufw allow in 443/tcp
ufw allow in 521/tcp
ufw logging off
ufw --force enable
ufw --force delete 4
ufw --force delete 4
ufw --force delete 4

# Optimize
systemctl disable unattended-upgrades.service apt-daily.timer apt-daily-upgrade.timer console-setup.service keyboard-setup.service remote-fs.target man-db.timer systemd-timesyncd.service
sed -i 's/MODULES=most/MODULES=dep/g' /etc/initramfs-tools/initramfs.conf
sed -i 's/COMPRESS=gzip/COMPRESS=lz4/g' /etc/initramfs-tools/initramfs.conf
update-initramfs -u
echo 'GRUB_TIMEOUT=0' >> /etc/default/grub
update-grub
apt-get -y purge apparmor exim\*

for i in $(seq 0 "$(nproc --ignore 1)"); do
  echo "devices/system/cpu/cpu${i}/cpufreq/scaling_governor = performance" >> /etc/sysfs.conf
done

# Disable sleep when closing laptop screen
echo HandleLidSwitch=ignore >> /etc/systemd/logind.conf

# noatime
sed -i 's| / ext4 | / ext4 noatime,|g' /etc/fstab

# Disable swap
swapoff -a
sed -i '/swap/d' /etc/fstab

# node.JS & yarn
curl -sL https://deb.nodesource.com/setup_12.x | bash -
apt-get -y install nodejs
curl -sL https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add -
echo "deb https://dl.yarnpkg.com/debian/ stable main" > /etc/apt/sources.list.d/yarn.list
apt-get update
apt-get -y install yarn

# Docker
apt-get -y install apt-transport-https ca-certificates curl gnupg2 software-properties-common
curl -fsSL https://download.docker.com/linux/debian/gpg | apt-key add -
add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/debian $(lsb_release -cs) stable"
apt-get update
apt-get -y install docker-ce docker-ce-cli containerd.io

# ripgrep
curl -LO https://github.com/BurntSushi/ripgrep/releases/download/12.0.1/ripgrep_12.0.1_amd64.deb
dpkg -i ripgrep_12.0.1_amd64.deb
rm ripgrep_12.0.1_amd64.deb

# Go
wget https://dl.google.com/go/go1.15.2.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.15.2.linux-amd64.tar.gz
rm go1.15.2.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> /etc/profile

# Netdata
bash <(curl -Ss https://my-netdata.io/kickstart-static64.sh) --no-updates --stable-channel --disable-telemetry --dont-wait

# Caddy
tmpdir=$(mktemp -d)
cd "$tmpdir"
wget https://github.com/caddyserver/caddy/releases/download/v1.0.4/caddy_v1.0.4_linux_amd64.tar.gz
tar -xf caddy_v1.0.4_linux_amd64.tar.gz
mv caddy /usr/local/bin
cd
rm -rf "$tmpdir"

# Generate SSH key
ssh-keygen -ted25519 -f ~/.ssh/id_ed25519 -N ''

# Cleanup
sed -i '/^deb-src/d' /etc/apt/sources.list
apt-get update
apt-get -y purge unattended-upgrades
apt-get -y autoremove --purge
apt-get clean

# The end
reboot
