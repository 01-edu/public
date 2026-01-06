#!/usr/bin/env bash

find . -type f -exec ls -lha {} \; | sort -hrk5 | head -7 | awk '{printf("%5s | ", $5); print $NF}'
