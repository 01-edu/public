## twosum

### Instructions

Given a slice of integers, return indexes of the two numbers such that they add up to a specific target.

If there are more than one solution, return the first one.

If there are no solutions, return nil.

Expected function :

```go
func TwoSum(nums []int, target int) []int {
}
```

Here is a possible program to test your function :

```go
package main

import "fmt"

func main() {
	case1 := []int{1, 2, 3, 4}
	out := TwoSum(case1, 5)
	fmt.Println(out)
}
```

And its output :

```console
$ go run .
[0 3]
$
```
