## count-negative

### Instructions 

Write a function that takes an array of integers and returns the number of negative numbers in the array.
- If the array is empty, the function should return `0`.

### Expected function 
```go 
func CountNegative(numbers []int) int {
    // your code here
}
```

### Usage 

Here is a possible program to test your function:

```go
package main 

import "fmt"

func main(){
    fmt.Println(CountNegative([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
    fmt.Println(CountNegative([]int{-1, -2, -3, -4, -5, -6, -7, -8, -9, -10}))
    fmt.Println(CountNegative([]int{}))
    fmt.Println(CountNegative([]int{-1,2,0,-3}))
}
```

and the output should be:

```console
$ go run . 
0
10
0
2
```
