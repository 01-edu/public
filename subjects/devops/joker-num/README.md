## joker-num

### Instructions

In this exercise you are going to create a guessing game. The script will prompt player one to submit a number between 1 and 100000 (inclusive) and hide the input from the other player. The script will use a for loop to give player two an infinite amount of chances to guess the number. You will also need to check if the input of player one is valid and if it is not ask again.

Player two will be prompted with a visible prompt to guess the number, then the script will check if the guess is correct, too low or too high.

Your script needs to check if the input is empty, not a number or not between 1 and 100000 (inclusive).

- If the input is empty the script will output `Error: Input is not a number` and continue the script execution.
- If the input contains multiple numbers, the script will output `Error: Input is not a number` and continue the script execution.
- If the input is not a number, the script will output as well `Error: Input is not a number`.
- If the input is not between 1 and 100000 (inclusive), the script will output `Error: Number out of range` and continue the script execution.

If the guess is correct the script will display `Congratulations! You guessed the number.` If the guess is lower than the number, the script will display `Go up.` If the guess is higher, the script will display `Go down.`

### Usage

```console
$ ./joker-num.sh
Player one, please enter a number between 1 and 100000 (inclusive) and press enter:
Player two, please enter your guess:
asdf
Error: Input is not a number
Player two, please enter your guess:
100
Go down.
Player two, please enter your guess:
30
Player two, please enter your guess:
34
Congratulations! You guessed the number.
```

### Hints

- `read` : this command can be used to read input from the command line. To be able to input and hide sensitive information like passwords, the `-s` flag can be used.

For example:

```console
$ read -s password
```

After running that line, you would type your password, press enter, and it would be stored in the `$password` variable for later use.
