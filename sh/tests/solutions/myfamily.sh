#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

f=""
# --insecure flag to make it work with dev environment (self-signed certificate)
test "$DOMAIN" = "localhost" && f="--insecure"

curl $f --compressed --silent --show-error --max-time 10 https://$DOMAIN/assets/superhero/all.json |
    jq ".[] | select(.id == $HERO_ID)" | grep relatives | cut -d'"' -f4
