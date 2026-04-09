#!/usr/bin/env bash
curl -s "https://$DOMAIN/assets/superhero/all.json" | jq -r '.[] | select(.id==170) | "\(.name)\n\(.powerstats.power)\n\(.appearance.gender)"'
