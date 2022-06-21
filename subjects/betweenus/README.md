## between-us

### Instructions

Write a function named `BetweenUs` that takes 3 paramters and return : 
- If the first paramter is between the second and third paramters, return **true** else return **false**
- If the second parameter is bigger than the third return **false**

### Expected function
```go
func BetweenUs(num, min, max int) bool {
    // Your code here
}
```


### Usage

Here is a possible program to test your function:

```go
package main

import "fmt"

func main(){
    fmt.Println(BetweenUs(1, 2, 3))
    fmt.Println(BetweenUs(1, 1, 3))
    fmt.Println(BetweenUs(1, 3, 3))
    fmt.Println(BetweenUs(1, 1, 1))
    fmt.Println(BetweenUs(1, 2, 1))
    fmt.Println(BetweenUs(-1, -10, 0))
}
```

and the output should be:

```console
$ go run .
false
true
false
true
false
true
```
