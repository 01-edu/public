#!/usr/bin/env bash

export MY_MESSAGE="Hello World"
export MY_NUM=100
export MY_PI=3.142
# Grep only new vars
printenv | grep "MY_"
