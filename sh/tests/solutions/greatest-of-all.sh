#!/usr/bin/env bash

largest=0

for i in {1..10}; do

    read -p "Enter a number: " num
    if [ $num -gt $largest ]; then
        largest=$num
    fi
done

echo "The largest number is: $largest"
