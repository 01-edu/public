## Write_a

### Instructions

Write a program that takes a string and displays the first 'a' Character it encounters in it, followed by a newline. 
- If there are no 'a' characters in the string, the program writes 'a' followed By a newline. 
- If the number of parameters is not 1, the program displays 'a' followed by a newline.

### Usage

```
$ go run . "abc" | cat -e
a$
$ go run . "welcom to 01" "talent" | cat -e
a$
$ go run . "have a nice day!" | cat -e
a$
$ go run . | cat -e
a$
```