#!/usr/bin/env bash

# Configure Terminal

script_dir="$(cd -P "$(dirname "$BASH_SOURCE")" && pwd)"
cd $script_dir
. set.sh

# Makes bash case-insensitive
cat <<EOF>> /etc/inputrc
set completion-ignore-case
set show-all-if-ambiguous On
set show-all-if-unmodified On
EOF

# Enhance Linux prompt
cat <<EOF> /etc/issue
Kernel build: \v
Kernel package: \r
Date: \d \t
IP address: \4
Terminal: \l@\n.\O

EOF

# Enable Bash completion
apt-get -y install bash-completion

cat <<EOF>> /etc/bash.bashrc
if ! shopt -oq posix; then
  if [ -f /usr/share/bash-completion/bash_completion ]; then
    . /usr/share/bash-completion/bash_completion
  elif [ -f /etc/bash_completion ]; then
    . /etc/bash_completion
  fi
fi
EOF

# Set-up all users
for dir in $(ls -1d /root /home/* 2>/dev/null ||:)
do
	# Hide login informations
	touch $dir/.hushlogin

	# Add convenient aliases & behaviors
	cat <<-'EOF'>> $dir/.bashrc
	export LS_OPTIONS="--color=auto"
	eval "`dircolors`"

	alias df="df --si"
	alias du="du -cs --si"
	alias free="free -h --si"
	alias l="ls $LS_OPTIONS -al --si"
	alias less="less -i"
	alias nano="nano -clDOST4"
	alias pstree="pstree -palU"

	GOPATH=$HOME/go
	HISTCONTROL=ignoreboth
	HISTFILESIZE=
	HISTSIZE=
	HISTTIMEFORMAT="%F %T "
	EOF

	# Fix rights
	usr=$(echo "$dir" | rev | cut -d/ -f1 | rev)
	chown -R $usr:$usr $dir ||:
done
