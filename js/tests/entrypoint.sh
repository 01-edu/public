#!/bin/sh

set -e

node /app/test.mjs "/jail/student" "${EXERCISE}"
