#!/usr/bin/env bash

# Run me with:
#
# bash <(curl -Ss raw.githubusercontent.com/01-edu/public/master/scripts/kickstart.sh)

wget github.com/01-edu/public/archive/master.zip
unzip master.zip
cd public-master/scripts
sudo ./install_client.sh
cat dconfig.txt | dconf load /
cd ../..
rm -rf master.zip public-master
reboot
