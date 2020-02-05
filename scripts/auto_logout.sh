#!/bin/bash

set -v -x

cat <<EOF> /etc/dconf/profile/user
user-db:user
system-db:local

EOF

mkdir /etc/dconf/db/local.d/

cat <<EOF> /etc/dconf/db/local.d/00-autologout
[org/gnome/settings-daemon/plugins/power]
# Set the timeout to 900 seconds when on mains power
sleep-inactive-ac-timeout=10
# Set action after timeout to be logout when on mains power
sleep-inactive-ac-type='logout'

EOF

mkdir /etc/dconf/db/local.d/locks

cat <<EOF>  /etc/dconf/db/local.d/locks/autologout
# Lock automatic logout settings
/org/gnome/settings-daemon/plugins/power/sleep-inactive-ac-timeout
/org/gnome/settings-daemon/plugins/power/sleep-inactive-ac-type

EOF

dconf update
