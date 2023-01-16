#!/bin/bash

# Loop for player one
for (( ; ; ))
do
    echo "Player one, please enter a number between 1 and 100000 (inclusive) and press enter:"
    read number
    # timeout 1s read -s number
    # read -p "Player one, please enter a number between 1 and 100000 (inclusive) and press enter:" -s number
    # sleep 5s
    # Check if input is empty
    if [[ -z "$number" ]]
    then
        echo "Error: Input is empty, please try again."

    # Check if input is a number
    elif ! [[ "$number" =~ ^[0-9]+$ ]]
    then
        echo "Error: Input is not a number, please try again."

    # Check if input is between 1 and 100000 (inclusive)
    elif [[ "$number" -lt 1 || "$number" -gt 100000 ]]
    then
        echo "Error: Number out of range, please try again."
      
    else
        break
    fi
done

# Start the for loop for player two
for (( ; ; ))
do
    echo "Player two, please enter your guess:"
    read guess
    if [[ -z "$guess" ]]
    then
        echo "Error: Input is empty"
        continue
    fi

    # Check if input is a number
    if ! [[ "$guess" =~ ^[0-9]+$ ]]
    then
        echo "Error: Input is not a number"
        continue
    fi

    # Check if guess is correct
    if [[ "$guess" -eq "$number" ]]
    then
        echo "Congratulations! You guessed the number."
        break
    fi

    # Check if guess is too low or too high
    if [[ "$guess" -lt "$number" ]]
    then
        echo "Go up."
    else
        echo "Go down."
    fi
done
