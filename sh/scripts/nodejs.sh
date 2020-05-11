#!/usr/bin/env bash

# Install Node.js

script_dir="$(cd -P "$(dirname "$BASH_SOURCE")" && pwd)"
cd $script_dir
. set.sh

curl -sL https://deb.nodesource.com/setup_12.x | bash -
apt-get -y install nodejs
