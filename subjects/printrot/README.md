## printrot

### Instructions

Write a program that takes a lower case letter and goes around in order to display the whole alphabet starting from
this letter on a single line.

A line is a sequence of characters preceding the [end of line](https://en.wikipedia.org/wiki/Newline) character (`'\n'`).

If the input is invalid the program prints a newline (`'\n'`).

### Usage

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test "abc"

student@ubuntu:~/[[ROOT]]/test$ ./test "a"
abcdefghijklmnopqrstuvwxyz
student@ubuntu:~/[[ROOT]]/test$ ./test "g"
ghijklmnopqrstuvwxyzabcdef
student@ubuntu:~/[[ROOT]]/test$ ./test "a" "a"

student@ubuntu:~/[[ROOT]]/test$ ./test "A"

student@ubuntu:~/[[ROOT]]/test$
```
