## lastword

### Instructions

Write a program that takes a `string` and displays its last word, followed by a newline (`'\n'`).

- A word is a section of `string` delimited by spaces or by the start/end of the `string`.

- The output will be followed by a newline (`'\n'`).

- If the number of arguments is different from 1, or if there are no word, the program displays nothing.

### Usage

```console
student@ubuntu:~/lastword$ go build
student@ubuntu:~/lastword$ ./lastword "FOR PONY" | cat -e
PONY$
student@ubuntu:~/lastword$ ./lastword "this        ...       is sparta, then again, maybe    not" | cat -e
not$
student@ubuntu:~/lastword$ ./lastword "  "
student@ubuntu:~/lastword$ ./lastword "a" "b"
student@ubuntu:~/lastword$ ./lastword "  lorem,ipsum  " | cat -e
lorem,ipsum$
student@ubuntu:~/lastword$
```
