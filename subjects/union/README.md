## union

### Instructions

Write a program that takes two `string` and displays, without doubles, the characters that appear in either one of the `string`.

The display will be in the same order that the characters appear on the command line and will be followed by a newline (`'\n'`).

If the number of arguments is different from 2, then the program displays a newline (`'\n'`).

### Usage

```console
$ go run . zpadinton paqefwtdjetyiytjneytjoeyjnejeyj | cat -e
zpadintoqefwjy$
$
$ go run . ddf6vewg64f gtwthgdwthdwfteewhrtag6h4ffdhsd | cat -e
df6vewg4thras$
$
$ go run . rien "cette phrase ne cache rien" | cat -e
rienct phas$
$
$ go run . | cat -e
$
$ go run . rien | cat -e
$
```
