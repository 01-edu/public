#!/bin/bash

# Configure Terminal

SCRIPT_DIR="$(cd -P "$(dirname "$BASH_SOURCE")" && pwd)"
cd $SCRIPT_DIR
. set.sh

# Makes bash case-insensitive
cat <<EOF>> /etc/inputrc
set completion-ignore-case On
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
for DIR in $(ls -1d /root /home/* 2>/dev/null || true)
do
	# Hide login informations
	touch $DIR/.hushlogin

	# Add convenient aliases & behaviors
	cat <<-'EOF'>> $DIR/.bashrc
	HISTCONTROL=ignoreboth
	export HISTFILESIZE=
	export HISTSIZE=
	export HISTTIMEFORMAT="%F %T "
	alias l="ls $LS_OPTIONS -al --si"
	alias df="df --si"
	alias du="du -cs --si"
	alias free="free -h --si"
	alias pstree="pstree -palU"
	EOF

	# Fix rights
	USR=$(echo "$DIR" | rev | cut -d/ -f1 | rev)
	chown -R $USR:$USR $DIR || true
done
