#!/usr/bin/env bash

# Setup everything

script_dir="$(cd -P "$(dirname "$BASH_SOURCE")" && pwd)"
cd $script_dir
. set.sh

repo_dir=$(git rev-parse --show-toplevel)

gsettings set org.gnome.desktop.session idle-delay 0
gsettings set org.gnome.desktop.screensaver lock-enabled false

sudo -E ./configure_ubuntu.sh
cat dconfig.txt | dconf load /
rm -rf "$repo_dir"
reboot
