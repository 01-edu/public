#!/usr/bin/env bash

set -euo pipefail
IFS='
'

if ! test -d "$HOME/VirtualBox VMs"; then
    echo "Is VirtualBox installed ? Could not find $HOME/VirtualBox VMs"
    exit 1
fi

echo "If you have created snapshots for the debian-01 virtual machine, they will be deleted"
echo "Same thing for display recordings"
echo "You can safely ignore this warning if you don't know what it means"
echo
echo "Press the [RETURN] key to continue"
read -r

if pgrep VirtualBox >/dev/null; then
    echo Close VirtualBox before updating debian-01 Virtual Machine
    exit 1
fi

key_file=$(mktemp)

cat <<'EOF'> "$key_file"
-----BEGIN OPENSSH PRIVATE KEY-----
b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW
QyNTUxOQAAACAeqfcTKSukPU9D1pmO2v7EeaTefBMlr84RAKfjR3xlwgAAAJDSLNAv0izQ
LwAAAAtzc2gtZWQyNTUxOQAAACAeqfcTKSukPU9D1pmO2v7EeaTefBMlr84RAKfjR3xlwg
AAAEB8CYX9ExfuZx3Uw7otY3CIa2G+5K71Zu3iPn/97LXInB6p9xMpK6Q9T0PWmY7a/sR5
pN58EyWvzhEAp+NHfGXCAAAADWFzc2V0c0BzZXJ2ZXI=
-----END OPENSSH PRIVATE KEY-----
EOF

rsync \
    --archive \
    --delete \
    --compress \
    --compress-choice lz4 \
    --info=progress2 \
    --rsh="ssh -oIdentitiesOnly=yes -p521 -i$key_file" \
    "$HOME/VirtualBox VMs/debian-01" assets@assets.01-edu.org:
    # assets@assets.01-edu.org:debian-01 "$HOME/VirtualBox VMs"

echo "debian-01 successfully updated"
