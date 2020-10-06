#!/bin/sh

set -e

cd /app
node --no-warnings --unhandled-rejections=strict test.js "${EXERCISE}"
