#!/usr/bin/env bash
echo WHAT
response=$(curl -s "https://$DOMAIN/api/graphql-engine/v1/graphql" --data '{"query":"{user(where:{login:{_eq:\"'$USERNAME'\"}}){id}}"}' | jq '.data.user[0].id')
echo $response
