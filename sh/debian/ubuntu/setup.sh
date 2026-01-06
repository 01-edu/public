#!/usr/bin/env bash

# Setup everything

set -euo pipefail
IFS='
'

# The value of this parameter is expanded like PS1 and the expanded value is the
# prompt printed before the command line is echoed when the -x option is set
# (see The Set Builtin). The first character of the expanded value is replicated
# multiple times, as necessary, to indicate multiple levels of indirection.
# \D{%F %T} prints date like this : 2019-12-31 23:59:59
PS4='-\D{%F %T} '

# Print commands and their arguments as they are executed.
set -x

script_dir=$(cd -P "$(dirname "${BASH_SOURCE[0]}")" && pwd)
cd "$script_dir"

repo_dir=$(git rev-parse --show-toplevel)

gsettings set org.gnome.desktop.session idle-delay 0
gsettings set org.gnome.desktop.screensaver lock-enabled false

sudo -E ./configure.sh
dconf load / <dconfig.txt
rm -rf "$repo_dir"
reboot
