## repeatalpha

### Instructions

Write a program called `repeat_alpha` that takes a `string` and displays it
repeating each alphabetical character as many times as its alphabetical index.

The result must be followed by a newline (`'\n'`).

`'a'` becomes `'a'`, `'b'` becomes `'bb'`, `'e'` becomes `'eeeee'`, etc...

The case remains unchanged.

If the number of arguments is different from 1, the program displays a newline (`'\n'`).

### Usage

```console
student@ubuntu:~/piscine-go/repeatalpha$ go build
student@ubuntu:~/piscine-go/repeatalpha$ ./repeatalpha "abc" | cat -e
abbccc
student@ubuntu:~/piscine-go/repeatalpha$ ./repeatalpha "Alex." | cat -e
Alllllllllllleeeeexxxxxxxxxxxxxxxxxxxxxxxx.$
student@ubuntu:~/piscine-go/repeatalpha$ ./repeatalpha "abacadaba 42!" | cat -e
abbacccaddddabba 42!$
student@ubuntu:~/piscine-go/repeatalpha$ ./repeatalpha | cat -e
$
student@ubuntu:~/piscine-go/repeatalpha$ ./repeatalpha "" | cat -e
$
student@ubuntu:~/piscine-go/repeatalpha$
```
