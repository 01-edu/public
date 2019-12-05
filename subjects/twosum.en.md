## twosum

### Instructions

Given an array of integers, return indexes of the two numbers such that they add up to a specific target.

If there are more than one solution, return the first one.

If there are no solutions, return nil.

Expected function:

```go
func TwoSum(nums []int, target int) []int {
}
```

Here is a possible [program](TODO-LINK) to test your function :

```go
package main

import (
	"fmt"
	piscine ".."
)

func main() {
	case1 := []int{1, 2, 3, 4}
	out := piscine.TwoSum(case1, 5)
	fmt.Println(out)
}
```

And its output :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
[0 3]
student@ubuntu:~/[[ROOT]]/test$
```
