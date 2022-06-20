## Devisor

### Instructions

Write a function that takes a positive  integer and returns the number of it's devisors.
- If the the number a is negative return 0.
- Test numbers from 0 to 99999999 .

### Expected function
```go
func Divisors(n int) int {        
    ...
}
```
### Usage

Here is a possible program to test your function:

```go
package main

import (
    "fmt"
    "piscine"
)

func main() {
    fmt.Println(piscine.Divisors(4))// 4 can be divided by 1 and 2 and 4
    fmt.Println(piscine.Divisors(5))//5 can be divided by 1 and 5
}

```

And its output :


```console
$ go run .
3 
2 
```