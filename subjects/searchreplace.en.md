## searchreplace

### Instructions

Write a program that takes 3 arguments, the first arguments is a string in which to replace a letter (2nd argument) by another one (3rd argument).

- If the number of arguments is not 3, just display a newline.

- If the second argument is not contained in the first one (the string) then the program simply rewrites the string followed by a newline.

### Expected output

```console
student@ubuntu:~/piscine/test$ go build
student@ubuntu:~/piscine/test$ ./test "hella there" "a" "o"
hello there
student@ubuntu:~/piscine/test$ ./test "abcd" "z" "l"
abcd
student@ubuntu:~/piscine/test$ ./test "something" "a" "o" "b" "c"

student@ubuntu:~/piscine/test$
```
