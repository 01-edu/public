## sumthemall

### Instructions

Create a program that receives a number of arguments representing whole numbers and displays the sum of the total. If there are not enough arguments or there is an overflow display 0 followed by a new line. If the argument is not a number display 0 followed by a new line.

### Usage

```console
$ go run . 23 34 53 | cat -e
110$
$ go run . 1000 2000 3000 4000 | cat -e
10000$
$ go run . a | cat -e
0$
```
