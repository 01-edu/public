#!/usr/bin/env bash
curl -s "https://assets.01-edu.org/superhero/all.json" | jq -r '.[] | select(.id==170) | "\(.name)\n\(.powerstats.power)\n\(.appearance.gender)"'
