#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

diff student/\"\\?\$*\'First_file\'*\$?\\\" solutions/\"\\?\$*\'First_file\'*\$?\\\"
