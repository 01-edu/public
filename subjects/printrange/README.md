## print-range 

### instructions

- Write a function named `PrintRange` that takes a start and end number and prints all the numbers in that range.
    - If the start number is greater than the end number, print the numbers in descending order.
    - If the start number is less than the end number, print the numbers in ascending order.
    - If the number is greater than `9` print just to `9`
    - If the number is less than `0` print just to `0`
    - If both numbers are less than `0` or greater than `9` print (`'\n'`)
    - the output should be space-separated and (`'\n'`) at the end.

### Function
```go
func PrintRange(start, end int) {
    // Your code here
}
```
### Usage
```go
package main

import "piscine"
import "fmt"

func main() {
    piscine.PrintRange(1, 10)
    piscine.PrintRange(10, 1)
    piscine.PrintRange(1, 1)
    piscine.PrintRange(10, 10)
    piscine.PrintRange(0, 9)
    piscine.PrintRange(-1, -10)
}
```

### Output
```console
$ go run . | cat -e
1 2 3 4 5 6 7 8 9$
9 8 7 6 5 4 3 2 1$
1$
$
0 1 2 3 4 5 6 7 8 9$
$
```
