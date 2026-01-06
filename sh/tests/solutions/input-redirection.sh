#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

echo "cat -e <<EOF
The current directory is: $PWD
The default paths are: $PATH
The current user is: $USERNAME
EOF
" >show-info.sh
