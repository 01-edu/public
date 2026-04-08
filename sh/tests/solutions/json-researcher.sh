#!/usr/bin/env bash
curl -s "https://assets.01-edu.org/superhero/all.json" | jq '.[] | select(.id==1)' | grep "name\|\"power\""
