#!/bin/bash

# Install FX: command-line JSON processing tool (https://github.com/antonmedv/fx)

SCRIPT_DIR="$(cd -P "$(dirname "$BASH_SOURCE")" && pwd)"
cd $SCRIPT_DIR
. set.sh

npm install -g fx
