#!/bin/bash

# Install Go

SCRIPT_DIR="$(cd -P "$(dirname "$BASH_SOURCE")" && pwd)"
cd $SCRIPT_DIR
. set.sh

apt-get -y install golang

# Set-up all users
for DIR in $(ls -1d /root /home/* 2>/dev/null || true)
do
	# Add convenient aliases & behaviors
	cat <<-'EOF'>> $DIR/.bashrc
	GOPATH=$HOME/go
	PATH=$PATH:$GOPATH/bin
	alias gobuild='CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w"'
	EOF
	echo 'GOPATH=$HOME/go' >> $DIR/.profile

	# Fix rights
	USR=$(echo "$DIR" | rev | cut -d/ -f1 | rev)
	chown -R $USR:$USR $DIR || true
done
