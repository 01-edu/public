## sumascii

### Instructions

Create a program that receives an argument and displays the byte sum in decimal format, of the ASCII value of each character. The sum will be a number representing the value after the sum of the bytes.

- If there is an invalid input print `0`.

### Usage

```console
$ go run . "hi" | cat -e
209$
$ go run . "something larger" | cat -e
107$
$ go run . "" | cat -e
0$
$ go run . "S" "A"
0$
$ go run . | cat -e
0$
```
