## descendcombn

### Instructions

- Write a function that prints all possible combinations of `n` different digits in descending order.

- n will be defined as : 10 > n > 0

Below are the references for the **printing format** expected.

- (for n = 1) '9, 8, ..., 3, 2, 1, 0'

- (for n = 3) '987, 986, ..., 020, 019, 018, 017, 016, 015, 014, 013, 012'

### Expected function

```go
func DescendCombN(n int) {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import "piscine"

func main() {
	piscine.DescendCombN(9)
	piscine.DescendCombN(3)
	piscine.DescendCombN(1)
}
```

And its output :

```console
$ go run .
9, 8, 7, 6, 5, 4, 3, 2, 1, 0
987, 986, ..., 018, 017, 016, 015, 014, 013, 012
987654321, ..., 976543210, 876543210
$
```
