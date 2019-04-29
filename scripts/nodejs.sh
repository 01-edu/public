#!/bin/bash

# Install Node.js

SCRIPT_DIR="$(cd -P "$(dirname "$BASH_SOURCE")" && pwd)"
cd $SCRIPT_DIR
. set.sh

curl -sL https://deb.nodesource.com/setup_10.x | bash -
apt-get -y install nodejs
