#!/usr/bin/env bash

# Install Sublime Text & Sublime Merge

script_dir="$(cd -P "$(dirname "$BASH_SOURCE")" && pwd)"
cd $script_dir
. set.sh

wget -qO - https://download.sublimetext.com/sublimehq-pub.gpg | apt-key add -
apt-get install -y apt-transport-https

cat <<EOF> /etc/apt/sources.list.d/sublime-text.list
deb https://download.sublimetext.com/ apt/stable/
EOF

apt-get update
apt-get install -y sublime-text sublime-merge libgtk2.0-0
