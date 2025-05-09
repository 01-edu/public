## Can Jump

Given an array of non-negative integers representing the number of steps you can take forward from each position, implement the function `CanJump()` which takes a slice of unsigned integers `[]uint` as input and returns a `boolean` value. This function should determine if it's possible to reach and stay at the last index of the array starting from the first index, based on the steps you need to advance. Be aware that:

- Each value represents the exact number of steps you must take forward from that position.
- The function should return `true` if it's possible to reach and stay at the last index without stepping out of the array, and `false` otherwise.
- If the input has only one element, that is the last position in the array so the function will return `true` but if the array is empty it returns `false`.

Let's take the example array input := []uint{2, 3, 1, 1, 4}.

```console

Position: 0  1  2  3  4
Steps:    2  3  1  1  4
          ^

// Starting from position 0, you have 2 steps to move forward. This means you will move to positions 2.

Position: 0  1  2  3  4
Steps:    2  3  1  1  4
                ^

// From position 2, you have 1 step, so you will move to position 3.

Position: 0  1  2  3  4
Steps:    2  3  1  1  4
                   ^

// Finally, from position 3, you have 1 step to reach the last index at position 4 confirming that it's possible so the output will be "True".

Position: 0  1  2  3  4
Steps:    2  3  1  1  4
                      ^
```

### Usage

Here is a possible program to test your function:

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Println(piscine.CanJump(input1))

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println(piscine.CanJump(input2))

	input3 := []uint{0}
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
