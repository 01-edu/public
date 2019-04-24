## rot13

### Instructions

Write a program that takes a string and displays it, replacing each of its
letters by the letter 13 spaces ahead in alphabetical order.

- 'z' becomes 'm' and 'Z' becomes 'M'. Case remains unaffected.

- The output will be followed by a newline.

- If the number of arguments is not 1, the program displays a newline.

### Expected output

```console
student@ubuntu:~/piscine/test$ go build
student@ubuntu:~/piscine/test$ ./test "abc"
nop
student@ubuntu:~/piscine/test$ ./test "hello there"
uryyb gurer
student@ubuntu:~/piscine/test$ ./test

student@ubuntu:~/piscine/test$
```
