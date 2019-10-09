## union

### Instructions

Write a program that takes two `string` and displays, without doubles, the characters that appear in either one of the `string`.

The display will be in the order characters appear in the command line, and will be followed by a newline (`'\n'`).

If the number of arguments is different from 2, the program displays newline (`'\n'`).

### Usage

```console
student@ubuntu:~/piscine-go/union$ go build
student@ubuntu:~/piscine-go/union$ ./union zpadinton "paqefwtdjetyiytjneytjoeyjnejeyj" | cat -e
zpadintoqefwjy$
student@ubuntu:~/piscine-go/union$ ./union ddf6vewg64f gtwthgdwthdwfteewhrtag6h4ffdhsd | cat -e
df6vewg4thras$
student@ubuntu:~/piscine-go/union$ ./union "rien" "cette phrase ne cache rien" | cat -e
rienct phas$
student@ubuntu:~/piscine-go/union$ ./union | cat -e
$
student@ubuntu:~/piscine-go/union$ ./union "rien" | cat -e
$
student@ubuntu:~/piscine-go/union$
```
