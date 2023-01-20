## joker-num

### Instructions

In this exercise you are going to create a guessing game.
Create the script `joker-num.sh` which will receive the number to guess as argument. The number should be between 1 and 100 (inclusive). The player will then have 5 tries to guess the number.
To achieve this you will have to use a `for` loop.

When guessing a valid number the output will be:

- Number bigger than the secret one: `Go down`.
- Number smaller than the secret one: `Go up`.
- Number equal than the secret one: `Congratulations, you found the number!`.
- Ran out of tries: `You lost, the number was <number>`.

### Usage

```console
$ ./joker-num.sh 55
Enter your guess (5 tries left):
50
Go up
Enter your guess (4 tries left):
75
Go down
Enter your guess (3 tries left):
55
Congratulations, you found the number in 3 moves!
$ ./joker-num.sh 70
Enter your guess (5 tries left):
70
Congratulations, you found the number in 1 moves!
$
```

### Error handling

- Wrong number of arguments or number not between `1` and `100`: The program will print `Error: wrong argument` and return `1`.
- The player writing something that is not a valid number as an answer: The program will ask again to enter a guess and the number of tries won't be decremented.

Here an example for the last scenario:

```console
$ ./joker-num.sh 100
Enter your guess (5 tries left):
aaaa
Enter your guess (5 tries left):
                                    # Here the player only press enter
Enter your guess (5 tries left):
10
Go up
Enter your guess (4 tries left):
20
Go up
Enter your guess (3 tries left):
10000
Enter your guess (3 tries left):
30
Go up
Enter your guess (2 tries left):
40
Go up
Enter your guess (1 tries left):
50
Go up
You lost, the number was 100
```

### Hints

- `read`: this command can be used to read input from the command line.

For example:

```console
$ read guess
```

After running that line, you would type your guess, press `Enter`, and it would be stored in the `$guess` variable for later use.

- `if [[ -z "$var" ]]`: could be helpful to check if a string is empty.
