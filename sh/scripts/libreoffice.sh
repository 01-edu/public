#!/usr/bin/env bash

# Install VSCode

script_dir="$(cd -P "$(dirname "$BASH_SOURCE")" && pwd)"
cd $script_dir
. set.sh

apt-get -y install libreoffice
