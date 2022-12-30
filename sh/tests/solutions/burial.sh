sleep 2 &
jobs -l | awk '{print $1, $3, $4, $5, $6}'
