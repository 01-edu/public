## printcombn

### Instructions

- Write a function that prints all possible combinations of **n** different digits in ascending order.

- n will be defined as : 0 < n < 10

Below are the references for the **printing format** expected.

- (for n = 1) '0, 1, 2, 3, ..., 8, 9'

- (for n = 3) '012, 013, 014, 015, 016, 017, 018, 019, 023,...689, 789'

### Expected function

```go
func PrintCombN(n int) {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import "piscine"

func main() {
	piscine.PrintCombN(1)
	piscine.PrintCombN(3)
	piscine.PrintCombN(9)
}
```

And its output :

```console
$ go run .
0, 1, 2, 3, 4, 5, 6, 7, 8, 9
012, 013, 014, 015, 016, 017, 018, ... 679, 689, 789
012345678, 012345679, ..., 123456789
$
```

> Be mindful of your program efficiency to avoid timeouts.

### Notions

- [01-edu/z01](https://github.com/01-edu/z01)
