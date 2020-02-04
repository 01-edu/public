#!/usr/bin/env bash

script_dir="$(cd -P "$(dirname "$BASH_SOURCE")" && pwd)"
cd $script_dir
. set.sh

apt-get install -y python3-pip

# Install package for student_logger service
pip3 install plyvel
