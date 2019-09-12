#!/usr/bin/env bash

# Run me with:
#
# bash <(curl -Ss raw.githubusercontent.com/01-edu/public/master/scripts/kickstart.sh)

gsettings set org.gnome.desktop.session idle-delay 0
gsettings set org.gnome.desktop.screensaver lock-enabled false
wget github.com/01-edu/public/archive/master.zip
unzip master.zip
cd public-master/scripts
sudo -E ./install_client.sh
cat dconfig.txt | dconf load /
cd ../..
rm -rf master.zip public-master
reboot
