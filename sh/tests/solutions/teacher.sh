#!/usr/bin/env bash

PWD=$(pwd)

for folder in $PWD
do
        cd "$folder" || exit
        INTERVIEWNUMBER=$(head -n 179 streets/Buckingham_Place | tail -n 1 | cut -d "#" -f2)
        echo "$INTERVIEWNUMBER"
        cat interviews/interview-"$INTERVIEWNUMBER"
        grep -A 4 L337 vehicles | grep -A 3 -B 1 Honda | grep -A 2 -B 2 Blue | grep -B 4 "Height: 6"
        cat memberships/AAA memberships/Delta_SkyMiles memberships/Museum_of_Bash_History memberships/Terminal_City_Library| grep "$MAIN_SUSPECT" | wc -l
done
