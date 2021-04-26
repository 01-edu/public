## firstword

### Instructions

Write a program that takes an argument and displays its first word, followed by a newline (`'\n'`).

- A word is a sequence of characters delimited by spaces or by the start/end of the argument.

- The output will be followed by a newline (`'\n'`).

- If the number of arguments is not 1, or if there are no words, the program displays nothing.

### Usage

```console
student@ubuntu:~/firstword/test$ go build
student@ubuntu:~/firstword/test$ ./test "hello there"
hello
student@ubuntu:~/firstword/test$ ./test "hello   .........  bye"
hello
student@ubuntu:~/firstword/test$ ./test
student@ubuntu:~/firstword/test$ ./test "hello" "there"
student@ubuntu:~/firstword/test$
```
