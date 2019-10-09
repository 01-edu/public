## searchreplace

### Instructions

Write a program that takes 3 arguments, the first argument is a `string` in which to replace a letter (the 2nd argument) by another one (the 3rd argument).

-   If the number of arguments is different from 3, the program displays a newline (`'\n'`).

-   If the second argument is not contained in the first one (the string) then the program rewrites the `string` followed by a newline (`'\n'`).

### Usage

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test "hella there" "a" "o"
hello there
student@ubuntu:~/piscine-go/test$ ./test "abcd" "z" "l"
abcd
student@ubuntu:~/piscine-go/test$ ./test "something" "a" "o" "b" "c"

student@ubuntu:~/piscine-go/test$
```
