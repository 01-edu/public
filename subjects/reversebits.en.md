## reversebits

### Instructions

Write a program that takes a `byte` in binary format and that reverses it `bit` by `bit` (like the
example)

### Expected output

```console
student@ubuntu:~/piscine-go/reversebits$ go build
student@ubuntu:~/piscine-go/reversebits$ ./reversebits
Not enough arguments.
student@ubuntu:~/piscine-go/reversebits$ ./reversebits 00100110
01100100
student@ubuntu:~/piscine-go/reversebits$ ./reversebits "djs"
The argument "djs" doesn't represent a byte
student@ubuntu:~/piscine-go/reversebits$ ./reversebits "0102039s"
The argument "0102039s" doesn't represent a byte
student@ubuntu:~/piscine-go/reversebits$
```
