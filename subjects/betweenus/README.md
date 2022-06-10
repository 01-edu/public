## between-us

write a function named `BetweenUs` that takes 3 paramters and checks if the first number is between the second and third number.
- If the first number is between the second and third number, return `true` else return `false`

### Function
```go
func BetweenUs(num, min, max int) bool {
    // Your code here
}
```


### Usage
```go
package main

import "fmt"

func main(){
    fmt.Println(BetweenUs(1, 2, 3))
    fmt.Println(BetweenUs(1, 1, 3))
    fmt.Println(BetweenUs(1, 3, 3))
    fmt.Println(BetweenUs(1, 1, 1))
    fmt.Println(BetweenUs(1, 2, 1))
}
```

```console
$ go run .
false
true
false
true
false
```
