## returna

### Instructions

Write a program that takes a `string` as an argument, if the `string` contains the letter `a` then print `Contains the letter` followed by a newline `\n`.

- If there is no character `a` in the `string`, the program writes `!(Contains the letter)` followed by a newline `\n`.
- If the number of parameters is not 1, the program displays a newline `\n`.

### Usage

```console
$ go run . "do I exist" | cat -e
!(Contains the letter)$
$
$ go run . "I do exist a" | cat -e
Contains the letter$
$
$ go run . "one" "two" | cat -e
$
```
