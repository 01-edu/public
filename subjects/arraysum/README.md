## array-sum

### Instructions

Write a function that takes an array of numbers and returns the sum of all the numbers in the array.
- If the array is empty, the function should return 0.

### Expected function

```go
func SumArray(numbers []int) int {
    // your code here
}
```

### Usage

Here is a possible program to test your function:

```go
package main

import (
    "fmt"
)

func main(){
    fmt.Println(SumArray([]int{1,2,3,4,5}))
    fmt.Println(SumArray([]int{}))
    fmt.Println(SumArray([]int{-1,-2,-3,-4,-5}))
    fmt.Println(SumArray([]int{-1,2,3,4,-5}))
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
