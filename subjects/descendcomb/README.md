## descendcomb

### Instructions

Write a function that prints in descending order and on a single line all possible combinations of two different two-digit numbers.

These combinations are separated by a comma and a space.

### Expected function

```go
func DescendComb() {

}
```

### Usage

Here is a possible program to test your function:

```go
package main

import "piscine"

func main() {
	piscine.DescendComb()
}
```

This is the incomplete output:

```console
$ go run . | cat -e
99 98, 99 97, 99 96, 99 95, 99 94, ..., 03 01, 03 00, 02 01, 02 00, 01 00$
```
