## ispowerof2

### Instructions

Write a program that determines if a given number is a power of 2. A number `n` is a power of 2 if it meets the following condition: **n = 2 <sup>m</sup>** where m is another integer number.

For example, **4 = 2 <sup>2</sup>** or **16 = 2 <sup>4</sup>** are power of 2 numbers while 24 is not.

This program must print `true` if the given number is a power of 2, otherwise it has to print `false`.

- If there is more than one or no argument, the program should print nothing.

- When there is only one argument, it will always be a positive valid `int`.

### Usage :

```console
$ go run . 2 | cat -e
true$
$ go run . 64
true
$ go run . 513
false
$ go run .
$ go run . 64 1024
$
```
