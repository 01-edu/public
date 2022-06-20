## add-if-positive 

### Instructions
Write a function that takes two numbers and adds them together if they are both positive.
- If either number is negative or one of them is negative and the other is positive return `0`.

### Expected function 
```go
func AddIfPositive(a int, b int) int {
    // your code here
}
```

### Usage

Here is a possible program to test your function:

```go
package main

import "fmt"

func main() {
    fmt.Println(AddIfPositive(1, 2))
    fmt.Println(AddIfPositive(1, -2))
    fmt.Println(AddIfPositive(-1, 2))
    fmt.Println(AddIfPositive(-1, -2))
    fmt.Println(AddIfPositive(10,20))
    fmt.Println(AddIfPositive(0,20))
}
```

and the output should be:

```console
$ go run .
3
0
0
0
30
20
```
