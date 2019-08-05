#!/bin/bash

# Install VSCode

SCRIPT_DIR="$(cd -P "$(dirname "$BASH_SOURCE")" && pwd)"
cd $SCRIPT_DIR
. set.sh

apt-get -y install libreoffice
