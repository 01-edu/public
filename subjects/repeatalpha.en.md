## repeatalpha

### Instructions

Write a program called `repeat_alpha` that takes a `string` and displays it repeating each alphabetical character as many times as its alphabetical index.

The result must be followed by a newline (`'\n'`).

`'a'` becomes `'a'`, `'b'` becomes `'bb'`, `'e'` becomes `'eeeee'`, etc...

If the number of arguments is different from 1, the program displays nothing.

### Usage

```console
student@ubuntu:~/[[ROOT]]/repeatalpha$ go build
student@ubuntu:~/[[ROOT]]/repeatalpha$ ./repeatalpha abc | cat -e
abbccc
student@ubuntu:~/[[ROOT]]/repeatalpha$ ./repeatalpha Choumi. | cat -e
CCChhhhhhhhooooooooooooooouuuuuuuuuuuuuuuuuuuuummmmmmmmmmmmmiiiiiiiii.$
student@ubuntu:~/[[ROOT]]/repeatalpha$ ./repeatalpha "abacadaba 01!" | cat -e
abbacccaddddabba 01!$
student@ubuntu:~/[[ROOT]]/repeatalpha$ ./repeatalpha
student@ubuntu:~/[[ROOT]]/repeatalpha$ ./repeatalpha "" | cat -e
$
student@ubuntu:~/[[ROOT]]/repeatalpha$
```
