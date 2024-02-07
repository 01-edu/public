## reverse-strings

### Instructions

Write a function that takes a slice of strings and returns a single string containing them in reverse order with a space between each element of the slice. If the slice is empty, return an empty string.

### Expected function

```go
func ReverseStrings(strs []string) string {

}
```

### Usage

Here is a possible program to test your function:

```go
package main

import "fmt"

func main(){
    fmt.Println(ReverseStrings([]string{"a", "b", "c"}))
    fmt.Println(ReverseStrings([]string{"Good","Morning!"}))
    fmt.Println(ReverseStrings([]string{"Hello World"}))
}
```

And its output :

```console
$ go run .
c b a
!gninroM dooG
dlroW olleH
```
