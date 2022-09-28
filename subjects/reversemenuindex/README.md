## reversemenuindex

### Instructions

The restaurant employees are having a really tough day and are delivering the customers' food in the wrong order. You need to fix the problem so that they can deliver it correctly.

Write a function `ReverseMenuIndex()` that takes a `slice` of strings as an argument and returns another `slice` of strings with the menu in the correct order.

- `append()` is not allowed for this exercise.

### Expected function

```go
func ReverseMenuIndex(menu []string) []string {

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
        fmt.Println(piscine.ReverseMenuIndex([]string{"desserts", "mains", "drinks", "starters"}))
}
```

And its output:

```console
$ go run . | cat -e
[starters drinks mains desserts]$
```
