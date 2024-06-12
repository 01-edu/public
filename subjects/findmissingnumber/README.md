## find-missing-number

### Instructions

Write a function that takes an array of numbers and returns the missing number.
- If the array is empty, return `-1`.
- If the array contains only one number or there is no missing number, return `-1`.
- if the array contains more than one missing number, return the minimum missing number.

### Expected function
```go
func FindMissingNumber(numbers []int) int {
    // Your code here
}
```

### Usage

Here is a possible program to test your function:

```go
package main

import "fmt"

func main(){
    fmt.Println(FindMissingNumber([]int{1, 2, 5,3, 6, 7, 8, 9, 10}))
    fmt.Println(FindMissingNumber([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}))
    fmt.Println(FindMissingNumber([]int{-10,12,32})
}
```

and the output should be:

```console
$ go run .
4
-1
-9
```
