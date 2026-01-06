#!/bin/bash

CLIENTID=$1
CLIENTSECRET=$2

TOKEN=$(curl -k -H "Content-Type: application/x-www-form-urlencoded" -H "Authorization: Basic $(echo -n ${CLIENTID}:${CLIENTSECRET} | base64)" --data "grant_type=client_credentials" "http://mydomain.com/auth/realms/myrealm/protocol/openid-connect/token" -s | jq -r .access_token)

echo $TOKEN
