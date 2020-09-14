#!/usr/bin/env bash

set -euo pipefail
IFS='
'

echo Installing common tools
sudo apt -y install man bash-completion git jq curl build-essential netcat wget psmisc lz4 file net-tools brotli unzip zip moreutils pv tree whois

echo Installing Go
wget https://golang.org/dl/go1.15.2.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.15.2.linux-amd64.tar.gz
rm go1.15.2.linux-amd64.tar.gz
cat <<'EOF'>> ~/.bashrc
GOPATH=$HOME/go
PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
EOF

echo Installing Go tools
GOPATH=$HOME/go
PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
go get golang.org/x/tools/cmd/goimports

echo Installing OpenSSH
sudo apt -y install openssh-client
ssh-keygen -t ed25519 -f ~/.ssh/id_ed25519 -N ''
echo "Copy the following line :"
cat ~/.ssh/id_ed25519.pub
