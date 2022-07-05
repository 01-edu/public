## print-range 

### Instructions

Write a function called `PrintRange` that given a range between two numbers, prints all numbers in that range.

- If the starting number is greater than the ending number, print the numbers in descending order, otherwise in ascending order.
- If the number is greater than `9` print only up to `9`
- If the number is less than `0` print only up to `0`
- If both numbers are less than `0` print (`'\n'`), the same applies when both numbers are greater than `9`.
- The output must be separated by spaces and (`'\n'`) at the end.

### Expected function

```go
func PrintRange(start, end int) {
    // Your code here
}
```

### Usage

Here is a possible program to test your function:

```go
package main

import "fmt"

func main() {
    PrintRange(1, 10)
    PrintRange(10, 1)
    PrintRange(1, 1)
    PrintRange(10, 10)
    PrintRange(0, 9)
    PrintRange(-1, -10)
}
```

and the output should be:

```console
$ go run . | cat -e
1 2 3 4 5 6 7 8 9$
9 8 7 6 5 4 3 2 1$
1$
$
0 1 2 3 4 5 6 7 8 9$
$
```
