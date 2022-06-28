## addfront

### Instructions

Write a function that takes a string and a slice of strings, this function will return a new slice of string with the given string prepended
- If the given string is empty you only need to return the given slice

### Expected function 

```go
func AddFront(s string, slice []string) []string {
    // your code here
}
```

### Usage

Here is a possible program to test your function:

```go
package main

import "fmt"

func main() {
    fmt.Println(AddFront("Hello", []string{"world"}))
    fmt.Println(AddFront("Hello", []string{"world", "!"}))
    fmt.Println(AddFront("Hello", []string{}))
}
```

and the output should be:

```console
$ go run .
[Hello world]
[Hello world !]
[Hello]
```
