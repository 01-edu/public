TZ=utc ls -l --time-style='+%F %R' | sed 1d | awk '{print $1, $6, $7, $8}'
