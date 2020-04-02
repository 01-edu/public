#!/usr/bin/env bash

curl -s https://$DOMAIN/assets/superhero/all.json | jq ".[] | select( .id == $HERO_ID)" | grep relatives | cut -d'"' -f4
