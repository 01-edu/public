#!/bin/bash

find . -iregex '.*\.\(txt\)' -exec basename {} \; | cut -d "." -f 1
