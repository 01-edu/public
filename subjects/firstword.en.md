## firstword

### Instructions

Write a program that takes a `string` and displays its first word, followed by a newline (`'\n'`).

- A word is a section of `string` delimited by spaces or by the start/end of the `string`.

- The output will be followed by a newline (`'\n'`).

- If the number of parameters is not 1, or if there are no words, the program displays a newline (`'\n'`).

### Usage

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test "hello there"
hello
student@ubuntu:~/[[ROOT]]/test$ ./test "hello   .........  bye"
hello
student@ubuntu:~/[[ROOT]]/test$ ./test

student@ubuntu:~/[[ROOT]]/test$ ./test "hello" "there"

student@ubuntu:~/[[ROOT]]/test$
```
