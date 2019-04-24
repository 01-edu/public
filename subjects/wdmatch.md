## switchcase

### Instructions

Write a program that takes two strings and checks whether it's possible to write the first string with characters from the second string, while respecting the order in which these characters appear in the second string.

- If it's possible, the program displays the string followed by a `\n`, otherwise it simply displays a `\n`.

- If the number of arguments is not 2, the program displays `\n`.

### Expected output

```console
student@ubuntu:~/piscine/test$ go build
student@ubuntu:~/piscine/test$ ./test "faya"  "fgvvfdxcacpolhyghbreda" 
faya
student@ubuntu:~/piscine/test$ ./test "faya" "fgvvfdxcacpolhyghbred"

student@ubuntu:~/piscine/test$ ./test "error" 

student@ubuntu:~/piscine/test$ ./test

student@ubuntu:~/piscine/test$
```
