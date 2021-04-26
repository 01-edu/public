#!/usr/bin/env bash

set -euo pipefail
IFS='
'

set -x

sudo apt -y install man bash-completion git jq curl build-essential netcat wget psmisc file net-tools brotli unzip zip moreutils pv tree whois openssh-client

wget https://golang.org/dl/go1.16.3.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.16.3.linux-amd64.tar.gz
rm go1.16.3.linux-amd64.tar.gz
cat <<'EOF'>> ~/.bashrc
GOPATH=$HOME/go
PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
EOF

GOPATH=$HOME/go
PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
go get -v golang.org/x/tools/cmd/goimports
go get -v github.com/01-edu/z01
go get -v github.com/stamblerre/gocode
mv "${GOPATH}/bin/gocode" "${GOPATH}/bin/gocode-gomod"
go get -v github.com/mdempsky/gocode
go get -v github.com/uudashr/gopkgs/v2/cmd/gopkgs
go get -v github.com/ramya-rao-a/go-outline
go get -v github.com/acroca/go-symbols
go get -v golang.org/x/tools/cmd/guru
go get -v golang.org/x/tools/cmd/gorename
go get -v github.com/cweill/gotests/...
go get -v github.com/fatih/gomodifytags
go get -v github.com/josharian/impl
go get -v github.com/davidrjenni/reftools/cmd/fillstruct
go get -v github.com/haya14busa/goplay/cmd/goplay
go get -v github.com/godoctor/godoctor
go get -v github.com/go-delve/delve/cmd/dlv
go get -v github.com/rogpeppe/godef
go get -v github.com/sqs/goreturns
go get -v golang.org/x/lint/golint
go get -v golang.org/x/tools/gopls
