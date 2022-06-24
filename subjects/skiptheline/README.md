## skip-the-line

### Instructions

Write a function that receives a string and a slice of strings. Return a new slice with the given string prepended.

### Expected function 

```go
func SkipTheLine(s string, slice []string) []string {
    // your code here
}
```

### Usage

Here is a possible program to test your function:

```go
package main

import "fmt"

func main() {
    fmt.Println(SkipTheLine("Hello", []string{"world"}))
    fmt.Println(SkipTheLine("Hello", []string{"","world", "!"}))
    fmt.Println(SkipTheLine("Hello", []string{}))
}
```

and the output should be:

```console
$ go run .
[Hello world]
[Hello world !]
[Hello]
```