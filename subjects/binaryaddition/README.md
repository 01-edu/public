## binary-addition

### Instructions

Write a function named `BinaryAddition(int,int)` that takes two integers and returns the sum of the two in binary in an array of `int`.
- If one of the integers is negative return `nil`
- Convert the argument to binary then add the two binary numbers together

### Expected function 
```go
func BinaryAddition(a int, b int) []int {
    // your code here
}
```

### Usage

Here is a possible program to test your function:

```go
package main

import "fmt"


func main(){
	fmt.Println(BinaryAddition(1, 1))
	fmt.Println(BinaryAddition(1, 2))
	fmt.Println(BinaryAddition(1, 3))
	fmt.Println(BinaryAddition(2, 1))
	fmt.Println(BinaryAddition(2, 2))
	fmt.Println(BinaryAddition(1, 16))
}
```

and the output should be:

```console
$ go run . 
[1 0]
[1 1]
[1 0 0]
[1 1]
[1 0 0]
[1 0 1]
[1 0 0 0 1]
```
