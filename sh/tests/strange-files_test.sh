#!/usr/bin/env bash

# Unofficial Bash Strict Mode
set -euo pipefail
IFS='
'

diff student/firstFile solutions/firstFile
diff student/\"medium_File\!\" solutions/\"medium_File\!\"
diff student/\"\\?\$*\'Hard_file\'*\$?\\\" solutions/\"\\?\$*\'Hard_file\'*\$?\\\"
