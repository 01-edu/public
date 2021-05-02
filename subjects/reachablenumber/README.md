## reachablenumber

### Instructions

Let us define a function f(x) by the following: first we add 1 to x, and then while the last digit of the number equals 0, we shall be deleting 0. Let us call 'y' reachable if we can apply **f** to **x** (zero or more times), and get **y**. 102 is reachable from 10098: f(f(f(10098))) = f(f(10099)) = f(101) = f(102). Any number is reachable from itself. You are given a positive number **n**, count how many integers are reachable from **n**.

### Expected function

```go
func ReachableNumber(n int) int {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import (
	"fmt"

	"piscine"
)

func main() {
	fmt.Println(piscine.ReachableNumber(1))
	fmt.Println(piscine.ReachableNumber(10))
	fmt.Println(piscine.ReachableNumber(1001))
}
```

And its output :

```console
$ go run .
9
19
36
$
```
