## rotatevowels

### Instructions

Write a **program** that checks the arguments for vowels.

- If the arguments contain vowels (for this exercise `y` is not considered a vowel) the program has to **"mirror"** the position of the vowels in the arguments (see the examples).
- If the number of arguments is less than 1, the program displays a new line ("`\n`").
- If the arguments do not have any vowels, the program just prints the arguments.

Example of output :

```console
$ go run . "Hello World" | cat -e
Hollo Werld$
$ go run . "HEllO World" "problem solved"
Hello Werld problom sOlvEd
$ go run . "str" "shh" "psst"
str shh psst
$ go run . "happy thoughts" "good luck"
huppy thooghts guod lack
$ go run . "aEi" "Ou"
uOi Ea
$ go run .

$
```
