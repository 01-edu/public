## displayz

### Instructions

Write a program that takes a `string`, and displays the first `z` character it encounters in it, followed by a newline (`'\n'`).
If there are no `z` characters in the string, the program just writes `z` followed by a newline (`'\n'`).
If the number of arguments is not 1, the program displays an `z` followed by a newline (`'\n'`).

### Usage

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test "xyz"
z
student@ubuntu:~/[[ROOT]]/test$ ./test "bcvbvZ"
z
student@ubuntu:~/[[ROOT]]/test$ ./test "nbv"
z
student@ubuntu:~/[[ROOT]]/test$ ./test
```
