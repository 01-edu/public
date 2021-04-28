## repeatalpha

### Instructions

Write a program called `repeat_alpha` that takes a `string` and displays it repeating each alphabetical character as many times as its alphabetical index.

The result must be followed by a newline (`'\n'`).

`'a'` becomes `'a'`, `'b'` becomes `'bb'`, `'e'` becomes `'eeeee'`, etc...

If the number of arguments is different from 1, the program displays nothing.

### Usage

```console
$ go run . abc | cat -e
abbccc
$ go run . Choumi. | cat -e
CCChhhhhhhhooooooooooooooouuuuuuuuuuuuuuuuuuuuummmmmmmmmmmmmiiiiiiiii.$
$ go run . "abacadaba 01!" | cat -e
abbacccaddddabba 01!$
$ go run .
$ go run . "" | cat -e
$
$
```
