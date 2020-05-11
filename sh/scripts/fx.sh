#!/usr/bin/env bash

# Install FX: command-line JSON processing tool (https://github.com/antonmedv/fx)

script_dir="$(cd -P "$(dirname "$BASH_SOURCE")" && pwd)"
cd $script_dir
. set.sh

npm install -g fx
