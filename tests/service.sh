#!/usr/bin/env bash

. "$WRAPPER"

check_command docker "https://docs.docker.com/install"

if test "$1" = "build"; then
	docker build go -t tests-go
	docker build sh -t tests-sh
fi
