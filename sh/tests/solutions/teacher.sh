#!/usr/bin/env bash

INTERVIEWNUMBER=$(head -n 179 streets/Buckingham_Place | tail -n 1 | cut -d "#" -f2)
echo $INTERVIEWNUMBER
cat interviews/interview-"$INTERVIEWNUMBER"
echo $MAIN_SUSPECT
