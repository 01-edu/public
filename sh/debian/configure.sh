#!/usr/bin/env bash

# Bash Strict Mode
set -euo pipefail
IFS='
'
cd -P "$(dirname "$0")"
PS4='-\D{%F %T} '

export DEBIAN_FRONTEND=noninteractive
export DEBIAN_PRIORITY=critical

# Fix Debian 10 bug (https://bugs.debian.org/cgi-bin/bugreport.cgi?bug=905409)
PATH=/sbin:/usr/sbin:$PATH

function sshKeyGen() {
    # Check if any of the keys already exist
    local existing_keys=()
    for key_type in all https runner; do
        if [[ -f ~/.ssh/ed25519_01edu_$key_type ]]; then
            existing_keys+=(~/".ssh/ed25519_01edu_$key_type")
        fi
    done

    # If any keys exist, warn the user and offer backup option
    if [[ ${#existing_keys[@]} -gt 0 ]]; then
        echo "Warning: The following SSH keys already exist:"
        printf "  %s\n" "${existing_keys[@]}"
        echo ""
        read -rp "Would you like to backup these keys and continue? (y/n): " choice

        if [[ $choice =~ ^[Yy]$ ]]; then
            # Create backup directory with timestamp
            backup_dir=~/.ssh/backup_$(date +%Y%m%d_%H%M%S)
            mkdir -p "$backup_dir"

            # Backup existing keys
            for key in "${existing_keys[@]}"; do
                key_name=$(basename "$key")
                cp "$key" "$backup_dir/"
                cp "$key.pub" "$backup_dir/" 2>/dev/null || true
                echo "Backed up $key to $backup_dir/$key_name"
            done
            echo "Keys backed up to $backup_dir"
        else
            echo "Operation cancelled. Existing keys were not modified."
            return 1
        fi
    fi

    # Create config directory if it doesn't exist
    mkdir -p ~/.ssh/config.d

    # Generate SSH key and create SSH config
    for key_type in all https runner; do
        ssh-keygen -t ed25519 -f ~/.ssh/ed25519_01edu_$key_type -N ''

        # Create SSH config for each key
        echo "Host github.com-01-edu-$key_type
        HostName github.com
        User git
        IdentityFile ~/.ssh/ed25519_01edu_$key_type" >~/.ssh/config.d/01-edu-$key_type.conf
    done

    # Include custom SSH configurations from the config directory if not already included
    if ! grep -q "Include ~/.ssh/config.d/*.conf" ~/.ssh/config; then
        echo "Include ~/.ssh/config.d/*.conf" >>~/.ssh/config
    fi
}

function sysConfig() {
    echo "Enter the server FQDN $(tput setaf 2)[System: $(hostname)]$(tput sgr0):"
    read -r serverFQDN
    hostnamectl set-hostname "$serverFQDN"

    echo "Enter the server Time Zone $(tput setaf 2)[System: $(cat /etc/timezone)]$(tput sgr0): "
    read -r serverTZ
    timedatectl set-timezone "$serverTZ"

    # Navigate to tmp
    cd /tmp

    # Debian stable OS
    apt-get update
    apt-get -y -o "Dpkg::Options::=--force-confdef" -o "Dpkg::Options::=--force-confold" upgrade
    apt-get -y dist-upgrade

    # Disable OpenStack SSH malware
    mv /home/debian/.ssh/authorized_keys /root/.ssh/authorized_keys || :
    sed -i '/Generated-by-Nova/d' /root/.ssh/authorized_keys || :
    chown root:root /root/.ssh/authorized_keys || :

    # Terminal goodies
    touch .hushlogin

    cat <<'EOF' >>/root/.bashrc
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

export HISTCONTROL=ignoreboth
export HISTFILESIZE=
export HISTSIZE=
export HISTTIMEFORMAT="%F %T "
export DOCKER_BUILDKIT=1 COMPOSE_DOCKER_CLI_BUILD=1
EOF

    cat <<'EOF' >>/etc/inputrc
set completion-ignore-case
set show-all-if-ambiguous On
set show-all-if-unmodified On
EOF

    cat <<'EOF' >>/etc/bash.bashrc
if ! shopt -oq posix; then
  if [ -f /usr/share/bash-completion/bash_completion ]; then
    . /usr/share/bash-completion/bash_completion
  elif [ -f /etc/bash_completion ]; then
    . /etc/bash_completion
  fi
fi
EOF

    # Basic packages
    apt-get -y install man bash-completion git ufw jq cron curl build-essential wget psmisc lz4 file net-tools brotli unzip zip moreutils dnsutils fail2ban xauth sysfsutils rsync iperf pv tree mc screen ssh iotop htop awscli whois sudo

    # Enable time synchronization
    timedatectl set-ntp true

    # Configure screen
    cat <<'EOF' >>/etc/screenrc
startup_message off
shell -$SHELL
defscrollback 100000
bind l eval clear "scrollback 0" "scrollback 100000"
EOF

    # Configure SSH
    cat <<'EOF' >>/etc/ssh/sshd_config
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
    ufw allow in 8080/tcp
    ufw allow in 8082/tcp
    ufw logging off

    # Optimize
    systemctl disable apt-daily.timer apt-daily-upgrade.timer remote-fs.target man-db.timer
    sed -i 's/MODULES=most/MODULES=dep/g' /etc/initramfs-tools/initramfs.conf
    sed -i 's/COMPRESS=gzip/COMPRESS=lz4/g' /etc/initramfs-tools/initramfs.conf
    echo 'RESUME=none' >>/etc/initramfs-tools/conf.d/resume
    update-initramfs -u
    echo 'GRUB_TIMEOUT=0' >>/etc/default/grub
    update-grub
    apt-get -y purge apparmor exim\*

    for i in $(seq 0 "$(nproc --ignore 1)"); do
        echo "devices/system/cpu/cpu${i}/cpufreq/scaling_governor = performance" >>/etc/sysfs.conf
    done

    # Disable sleep when closing laptop screen
    echo HandleLidSwitch=ignore >>/etc/systemd/logind.conf

    # noatime
    sed -i 's| / ext4 | / ext4 noatime,|g' /etc/fstab

    # Disable swap
    swapoff -a
    sed -i '/swap/d' /etc/fstab

    # Docker
    curl -fsSL https://get.docker.com -o get-docker.sh && sh get-docker.sh

    # NodeJS
    curl -fsSL https://deb.nodesource.com/setup_22.x | bash && apt-get install -y nodejs

    # Git
    apt-get update && apt-get -y install git

    # Create the config.d directory if it doesn't exist
    mkdir -p ~/.ssh/config.d

    sshKeyGen

    # Use Cloudflare DNS server
    echo 'supersede domain-name-servers 1.1.1.1;' >>/etc/dhcp/dhclient.conf

    # Cleanup
    sed -i '/^deb-src/d' /etc/apt/sources.list
    apt-get update
    apt-get -y purge unattended-upgrades
    apt-get -y autoremove --purge
    apt-get clean

    # SSH Keys Infra Team
    curl https://github.com/{harryvasanth,frenchris,kigiri}.keys >>~/.ssh/authorized_keys

    # Create Core directories
    mkdir -p /root/core/scripts/misc

}

# Check Config
function checkConfig() {
    test "$(command -v "${1:-}")" && echo -n ✅ || echo -n ❌
    echo " $*"
}

function checkKeys() {
    # Check if SSH key pairs are generated
    if test -f ~/.ssh/ed25519_01edu_all && test -f ~/.ssh/ed25519_01edu_all.pub &&
        test -f ~/.ssh/ed25519_01edu_https && test -f ~/.ssh/ed25519_01edu_https.pub &&
        test -f ~/.ssh/ed25519_01edu_runner && test -f ~/.ssh/ed25519_01edu_runner.pub; then
        echo "✅ SSH private/public key pairs generated"

        # Echo public keys
        echo -e "$(tput setaf 2)$(tput bold)\n🔑 Public keys:$(tput sgr0)"
        echo "all"
        cat ~/.ssh/ed25519_01edu_all.pub
        echo "https"
        cat ~/.ssh/ed25519_01edu_https.pub
        echo "runner"
        cat ~/.ssh/ed25519_01edu_runner.pub
    else
        echo "❌ SSH private/public key pairs not found"
    fi
}

# Check configs in the List
function checkList() {
    checkConfig docker
    checkConfig node
    checkConfig git
    checkConfig man
    checkConfig ufw
    checkConfig jq
    checkConfig cron
    checkConfig curl
    checkConfig wget
    checkConfig lz4
    checkConfig file
    checkConfig brotli
    checkConfig unzip
    checkConfig zip
    checkConfig fail2ban-server
    checkConfig xauth
    checkConfig rsync
    checkConfig iperf
    checkConfig pv
    checkConfig tree
    checkConfig mc
    checkConfig screen
    checkConfig ssh
    checkConfig iotop
    checkConfig htop
    checkConfig aws
    checkConfig whois
    checkConfig sudo

    checkKeys
}

function runHTTPS() {
    echo -e "Deploying HTTPS service: \n"
    echo "Enter the server FQDN $(tput setaf 2)[System: $(hostname)]$(tput sgr0):"
    read -r httpsFQDN
    # Check if the FQDN is valid
    if dig +short "$httpsFQDN" >/dev/null 2>&1; then
        cd /root/core/https
        DOMAIN=$httpsFQDN ./run.sh
        echo -e "HTTPS service is up! \n"
    else
        echo "$(tput setaf 1)$(tput bold)The FQDN: $httpsFQDN is not reachable$(tput sgr0)"
        echo "$(tput setaf 1)Please check your DNS configuration$(tput sgr0)"
        runHTTPS
    fi
}

# Deploy core repositories
function deployCore() {
    # Check for the presence of configurations
    test "$(ls ~/.ssh/*.pub 2>/dev/null)" && echo -n "$(tput setaf 2)$(tput bold)Config check passed!$(tput sgr0)" || exit 1
    echo -e "$(tput setaf 6)$(tput bold)\nThe core components will be deployed to the server: $(tput sgr0)\n"

    # Clone core repositories
    git clone git@github.com-01-edu-runner:01-edu/runner.git /root/core/runner
    git clone git@github.com-01-edu-https:01-edu/https.git /root/core/https

    # Docker login
    echo -e "Enter the docker username: "
    read -r dockerUsername
    echo -e "Enter the docker password: "
    read -r dockerPassword
    docker login docker.01-edu.org -u "$dockerUsername" -p "$dockerPassword"

    # Deploy HTTPS
    runHTTPS

    # Deploy Runner
    echo -e "Deploying Runner service: \n "
    cd /root/core/runner

    # Get the latest release version tag and create latest branch
    tag=$(git describe --tags "$(git rev-list --tags --max-count=1)")
    git checkout "$tag"

    # Get user auth infor for the runner
    echo -e "Enter the runner Registry password: "
    read -r registryPassword
    echo -e "Enter the runner GitHub username: "
    read -r githubUsername
    echo -e "Enter the runner GitHub token: "
    read -r githubToken
    REGISTRY_PASSWORD=$registryPassword GITHUB_USERNAME=$githubUsername GITHUB_TOKEN=$githubToken ./run.sh
    echo -e "Runner service is up! \n"
}

function deployPlatform() {
    # Check for the presence of configurations
    test "$(ls ~/.ssh/*.pub 2>/dev/null)" && echo -n "$(tput setaf 2)$(tput bold)Config check passed!$(tput sgr0)" || exit 1
    echo -e "$(tput setaf 6)$(tput bold)\nThe platform components will be deployed to the server: $(tput sgr0)\n"

    # Clone platform repository
    echo "Enter the server FQDN $(tput setaf 2)[System: $(hostname)]$(tput sgr0):"
    read -r serverFQDN
    git clone git@github.com-01-edu-all:01-edu/all.git /root/"$serverFQDN"
    cd /root/"$serverFQDN"
    # Generate platform environment file automatically
    ./gen-env.sh.sh --gen
    ./redeploy.sh --hard
}

function clonePlatform() {
    # Check for the presence of configurations
    test "$(ls ~/.ssh/*.pub 2>/dev/null)" && echo -n "$(tput setaf 2)$(tput bold)Config check passed!$(tput sgr0)" || exit 1
    echo -e "$(tput setaf 6)$(tput bold)\nThe platform components will be deployed to the server: $(tput sgr0)\n"

    # Clone platform repository
    echo "Enter the target directory for the platform $(tput setaf 2)[System: $(pwd)/$(hostname)]$(tput sgr0):"
    read -r serverDir
    git clone git@github.com:01-edu/all.git "$serverDir"
    cd "$serverDir"
    # Generate platform environment file automatically
    ./generate_env.sh --gen
}

if [[ -z ${1:-} ]] || [[ "--check" = "$1" ]]; then
    echo -e "$(tput setaf 2)$(tput bold)Commencing configuration check: $(tput sgr0)"
    checkList
    echo -e "$(tput setaf 2)\nSystem configuration check complete! $(tput sgr0)\n"
    exit 0
elif [[ "--help" = "$1" ]]; then
    echo "$(tput setaf 2) --check : to check the current configuration. $(tput sgr0)"
    echo "$(tput setaf 3) --run : to configure the system. $(tput sgr0)"
    echo "$(tput setaf 1) --reboot : to configure the system and reboot. $(tput sgr0)"
    echo "$(tput setaf 6) --deploy : to deploy and spin-up platform components. $(tput sgr0)"
    echo "$(tput setaf 5) --ssh-keys : to generate ssh-key pairs. $(tput sgr0)"
    echo "$(tput setaf 5) --platform : to clone platform. $(tput sgr0)"
    echo "$(tput setaf 7) --help : to display this message. $(tput sgr0)"
elif [[ "--reboot" = "$1" ]]; then
    echo -e "$(tput setaf 1)$(tput bold)\nSystem will be configured and rebooted. $(tput sgr0)"
    sysConfig
    echo -e "$(tput setaf 1)\nSystem configuration complete. Rebooting now... $(tput sgr0)"
    reboot
elif [[ "--run" = "$1" ]]; then
    echo -e "$(tput setaf 3)$(tput bold)\nSystem will be configured without rebooting. $(tput sgr0)"
    sysConfig
    echo -e "$(tput setaf 3)\nSystem configuration complete! $(tput sgr0)"
    exit 0
elif [[ "--deploy" = "$1" ]]; then
    deployCore
    deployPlatform
    echo -e "$(tput setaf 6)\nRepositories cloned and platform has been deployed successfully! $(tput sgr0)"
    exit 0
elif [[ "--ssh-keys" = "$1" ]]; then
    sshKeyGen
    echo -e "$(tput setaf 5)\nSSH keys have been generated successfully! $(tput sgr0)"
    checkKeys
    exit 0
elif [[ "--deploy-core" = "$1" ]]; then
    deployCore
    echo -e "$(tput setaf 5)\nRepositories cloned and core has been deployed successfully! $(tput sgr0)"
    exit 0
elif [[ "--platform" = "$1" ]]; then
    clonePlatform
    echo -e "$(tput setaf 5)\nPlatform has been cloned successfully! $(tput sgr0)"
    exit 0
else
    echo "$(tput setaf 1)$(tput bold) Unknown configuration option: $1 $(tput sgr0)"
    echo "$(tput setaf 1)Please use --help for all available options. $(tput sgr0)"
    echo "$(tput setaf 1)No changes are made $(tput sgr0)"
    exit 0
fi
