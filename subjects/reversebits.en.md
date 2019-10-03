## reversebits

### Instructions

Write a program that takes a `byte` in binary format, that reverses it `bit` by `bit` (as shown in the
example) and that prints the result.

### Expected output

```console
student@ubuntu:~/piscine-go/reversebits$ go build
student@ubuntu:~/piscine-go/reversebits$ ./reversebits
Not enough arguments.
student@ubuntu:~/piscine-go/reversebits$ ./reversebits 00100110 | cat -e
01100100$
student@ubuntu:~/piscine-go/reversebits$ ./reversebits "djs"
The argument "djs" does not represent a byte
student@ubuntu:~/piscine-go/reversebits$ ./reversebits "0102039s" | cat -e
The argument "0102039s" does not represent a byte$
student@ubuntu:~/piscine-go/reversebits$
```
