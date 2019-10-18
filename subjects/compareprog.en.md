## compareprog

### Instructions

Write a program that behaves like the `Compare` function.

This program prints a number after comparing two `string` lexicographically.

### Usage :

```console
student@ubuntu:~/compareprog$ go build
student@ubuntu:~/compareprog$ ./compareprog a b | cat -e
-1$
student@ubuntu:~/compareprog$ ./compareprog a a | cat -e
0$
student@ubuntu:~/compareprog$ ./compareprog b a | cat -e
1$
student@ubuntu:~/compareprog$ ./compareprog b a d | cat -e
$
student@ubuntu:~/compareprog$ ./compareprog | cat -e
$
student@ubuntu:~/compareprog$
```
