## remove-odd

### Instructions

Write a function named `RemoveOdd(string)` that takes a string and returns a new string with all the odd characters removed.
- The function should skip spaces.

### Expected function 
```go
func RemoveOdd(string) string{
    // your code here
}
```

### Usage

``` go
package main

import "fmt"

func main(){
    fmt.Println(RemoveOdd("Hello World"))
    fmt.Println(RemoveOdd("H"))
    fmt.Println(RemoveOdd("How are you?"))
}
```

and the output should be:

```console
$ go run . 
Hlo Wrd
H
Hw ae yu
```