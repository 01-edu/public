## atoibase

### Instructions

Write a program that takes a `string` number and its `string` base as arguments and prints its conversion as an `int`.

- If the base or the `string` number is not valid it returns `0`.

- If the number of arguments is bigger or lower that two it should print a newline `\n`.

Validity rules for a base :

- A base must contain at least 2 characters.
- Each character of a base must be unique.
- A base should not contain `+` or `-` characters.

Only valid `string` numbers will be tested.

The program **does not have** to manage negative numbers.

### Expected output :

```console
student@ubuntu:~/atoibaseprog$ go build
student@ubuntu:~/atoibaseprog$ ./atoibaseprog 125 0123456789
125
student@ubuntu:~/atoibaseprog$ ./atoibaseprog 1111101 01
125
student@ubuntu:~/atoibaseprog$ ./atoibaseprog 7D 0123456789ABCDEF
125
student@ubuntu:~/atoibaseprog$ ./atoibaseprog uoi choumi
125
student@ubuntu:~/atoibaseprog$ ./atoibaseprog bbbbbab -ab
0
student@ubuntu:~/atoibaseprog$ ./atoibaseprog 1111101

student@ubuntu:~/atoibaseprog$
```
