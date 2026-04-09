#!/usr/bin/env bash
curl -s "https://$DOMAIN/assets/superhero/all.json" | jq '.[] | select(.id==1)' | grep "name\|\"power\""
