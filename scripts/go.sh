#!/usr/bin/env bash

# Install Go

script_dir="$(cd -P "$(dirname "$BASH_SOURCE")" && pwd)"
cd $script_dir
. set.sh

apt-get -y install golang

# Set-up all users
for dir in $(ls -1d /root /home/* 2>/dev/null ||:)
do
	# Add convenient aliases & behaviors
	cat <<-'EOF'>> $dir/.bashrc
	GOPATH=$HOME/go
	PATH=$PATH:$GOPATH/bin
	alias gobuild='CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w"'
	EOF
	echo 'GOPATH=$HOME/go' >> $dir/.profile

	# Fix rights
	usr=$(echo "$dir" | rev | cut -d/ -f1 | rev)
	chown -R $usr:$usr $dir ||:
done
