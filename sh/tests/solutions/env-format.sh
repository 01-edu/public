printenv PWD
printenv | grep "H" | awk -F "=" '{print $1}'
