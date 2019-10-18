## itoabaseprog

### Instructions

Write a program that:

- converts a number in base 10 into the number in the specified base
- receives two parameters:
  - the first is the value
  - the second is the base

The base is expressed as an `int`, from 2 to 16. The characters comprising
the base are the digits from 0 to 9, followed by uppercase letters from A to F.

For example, the base `4` would be the equivalent of "0123" and the base `16` would be the equivalent of "0123456789ABCDEF".

If the value is negative, the resulting `string` has to be preceded with a
minus sign `-`.

### Expected output

```console
student@ubuntu:~/piscine-go/itoabaseprog$ go build
student@ubuntu:~/piscine-go/itoabaseprog$ ./itoabaseprog 15 16
F
student@ubuntu:~/piscine-go/itoabaseprog$ ./itoabaseprog 255 2
11111111
student@ubuntu:~/piscine-go/itoabaseprog$ ./itoabaseprog -4 2 | cat -e
-100$
student@ubuntu:~/piscine-go/itoabaseprog$ ./itoabaseprog -df sdf
The value "-df" can not be converted to an int
student@ubuntu:~/piscine-go/itoabaseprog$ ./itoabaseprog 23 ew | cat -e
The value "ew" can not be converted to an int$
student@ubuntu:~/piscine-go/itoabaseprog$ ./itoabaseprog 4 23

student@ubuntu:~/piscine-go/itoabaseprog$ 323 12
22B
student@ubuntu:~/piscine-go/itoabaseprog$
```
