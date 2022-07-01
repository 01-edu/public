## concatenate

### Instructions

Write a program that displays all the arguments as a single string by simply putting one after each other. If there are not enough arguments it should display a new line.

### Usage

```console
$ go run . hello there chris | cat -e
hellotherechris$
$ go run . a b c d | cat -e
abcd$
$ go run . "hello there" "how are you?"
hello therehow are you?$
$ go run . "hello there" " " | cat -e
hello there $
```
