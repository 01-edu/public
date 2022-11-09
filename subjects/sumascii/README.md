## sumascii

### Instructions

Create a program that receives an argument and displays the sum of the ASCII value of each character. The sum will be a number representing the value after the sum of the bytes. 

- If there is an invalid input print `0`.

### Usage

```console
$ go run . "hi" | cat -e
209$
$ go run . "" | cat -e
0$
$ go run . "S" "A"
0$
$ go run . | cat -e
0$
```
