## capitalizeprog

### Instructions

Write a program that capitalizes the first letter of each word **and** lowercases the rest of each word of a `string`.

-   A word is a sequence of **alphanumerical** characters.

-   If there is more than one argument the program should print `Too many arguments`.

-   If there is no arguments given the program should print a newline ("`\n`").

### Usage :

```console
student@ubuntu:~/capitalizeprog$ go build
student@ubuntu:~/capitalizeprog$ ./capitalizeprog "Hello! How are you? How+are+things+4you?" | cat -e
Hello! How Are You? How+Are+Things+4you?$
student@ubuntu:~/capitalizeprog$ ./capitalizeprog Hello! How are you? | cat -e
Too many arguments$
student@ubuntu:~/capitalizeprog$ ./capitalizeprog

student@ubuntu:~/capitalizeprog$
```
