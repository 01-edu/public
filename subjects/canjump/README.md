### Can Jump

Given an array of integers representing the maximum number of steps you can take forward from each position, implement the function `CanJump()` which takes an `integer slice` as input and returns a `boolean` value to determine if it's possible to reach the last index starting from the first index based on these maximum steps. The function should return `true` if it's possible to reach the last index and `false` otherwise.

> Note: The function only needs to consider positive numbers or zero in the array of steps. Also remember if the input has only one element that is the last position in the array so the function will return `true`.

### Usage

Here is a possible program to test your function:

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	input1 := []int{2, 3, 1, 1, 4}
	fmt.Println(piscine.CanJump(input1))

	input2 := []int{3, 2, 1, 0, 4}
	fmt.Println(piscine.CanJump(input2))

	input3 := []int{0}
	fmt.Println(piscine.CanJump(input3))
}
```

And its output :

```console
$ go run .
true
false
true
$
```
