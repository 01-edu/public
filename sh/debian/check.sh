#!/usr/bin/env bash

set -euo pipefail
IFS='
'

dir=$(cd -P "$(dirname "$BASH_SOURCE")" && pwd)

check() {
	test "$(command -v "$1")" && echo -n ✅ || echo -n ❌
	echo " $@"
}

check go '(see : https://golang.org/dl & https://golang.org/doc/install)'
check gofmt '(see : https://golang.org/dl & https://golang.org/doc/install)'
check goimports '(run : go get golang.org/x/tools/cmd/goimports)'
check git '(see : https://git-scm.com/downloads)'
check jq '(see : https://stedolan.github.io/jq/download)'
check sed
check tar
check find
check cut
check awk
check grep
check wget
check curl
check diff
check chmod
check chown
check touch
check wc
check cat

test "$(ls ~/.ssh/*.pub 2>/dev/null)" && echo -n ✅ || echo -n ❌
echo " SSH public key (run : ssh-keygen -t ed25519 -f ~/.ssh/id_ed25519 -N '')"

# git repository configured
