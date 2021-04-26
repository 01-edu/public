## printrot

### Instructions

Write a program that takes a lower case letter and goes around in order to display the whole alphabet starting from
this letter on a single line.

A line is a sequence of characters preceding the [end of line](https://en.wikipedia.org/wiki/Newline) character (`'\n'`).

If the input is invalid the program prints a newline (`'\n'`).

### Usage

```console
student@ubuntu:~/printrot/test$ go build
student@ubuntu:~/printrot/test$ ./test "abc"

student@ubuntu:~/printrot/test$ ./test "a"
abcdefghijklmnopqrstuvwxyz
student@ubuntu:~/printrot/test$ ./test "g"
ghijklmnopqrstuvwxyzabcdef
student@ubuntu:~/printrot/test$ ./test "a" "a"

student@ubuntu:~/printrot/test$ ./test "A"

student@ubuntu:~/printrot/test$
```
