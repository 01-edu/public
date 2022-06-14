## array-sum

### Instructions

Write a function that takes an array of numbers and returns the sum of all the numbers in the array.
- If the array is empty, the function should return 0.

### Expected function

```go
func ArraySum(numbers []int) int {
    // your code here
}
```

### Usage

```go
package main

import (
    "fmt"
)

func main(){
    fmt.Println(ArraySum([]int{1,2,3,4,5}))
    fmt.Println(ArraySum([]int{}))
    fmt.Println(ArraySum([]int{-1,-2,-3,-4,-5}))
    fmt.Println(ArraySum([]int{-1,2,3,4,-5}))
}
```

and the output should be:

```console
$ go run . 
15
0
-15
3
```