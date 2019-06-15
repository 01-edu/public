## lastword

### Instructions

Write a program that takes a `string` and displays its last word, followed by a newline(`'\n'`).

- A word is a section of `string` delimited by spaces or by the start/end of the `string`.

- The output will be followed by a newline(`'\n'`).

- If the number of parameters is not 1, or if there are no words, the program displays a newline(`'\n'`).

### Usage

```console
student@ubuntu:~/student/lastword$ go build
student@ubuntu:~/student/lastword$ ./lastword "FOR PONY" | cat -e
PONY$
student@ubuntu:~/student/lastword$ ./lastword "this        ...       is sparta, then again, maybe    not" | cat -e
not$
student@ubuntu:~/student/lastword$ ./lastword "  " | cat -e
$
student@ubuntu:~/student/lastword$ ./lastword "a" "b" | cat -e
$
student@ubuntu:~/student/lastword$ ./lastword "  lorem,ipsum  " | cat -e
lorem,ipsum$
student@ubuntu:~/student/lastword$
```
