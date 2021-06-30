#!/usr/bin/env bash

echo "USERNAME: $USERNAME"
curl -s "https://$DOMAIN/api/graphql-engine/v1/graphql" --data '{"query":"{user(where:{githubLogin:{_eq:\"'$USERNAME'\"}}){id}}"}' | jq '.data.user[0].id'
