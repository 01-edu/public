## pingpong

### Instructions

Write a program that checks if all numbers up to a given number are odd or even.
The numbers will be between 0 and 100.

- If the number is odd then it prints the number and `pong` as in the example below.
- If the number is even then it prints the number and `ping` in the same way.
- If the number of arguments is invalid then do nothing.

### Usage

```console
$ go run . 6 | cat -e
0 ping$
1 pong$
2 ping$
3 pong$
4 ping$
5 pong$
$ go run . | cat -e
$
```
