## firstword

### Instructions

Write a program that takes a string and displays its first word, followed by a newline.

- A word is a section of string delimited by spaces or by the start/end of the string.

- The output will be followed by a newline.

- If the number of parameters is not 1, or if there are no words, simply display a newline.

### Expected output

```console
student@ubuntu:~/piscine/test$ go build
student@ubuntu:~/piscine/test$ ./test "hello there"
hello
student@ubuntu:~/piscine/test$ ./test "hello   .........  bye"
hello
student@ubuntu:~/piscine/test$ ./test

student@ubuntu:~/piscine/test$ ./test "hello" "there"

student@ubuntu:~/piscine/test$
```
