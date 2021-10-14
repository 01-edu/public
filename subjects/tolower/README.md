## tolower

### Instructions

Write a function that lower cases for each letter in a `string`.

### Expected function

```go
func ToLower(s string) string {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.ToLower("Hello! How are you?"))
}
```

And its output :

```console
$ go run .
hello! how are you?
$
```

### Notions

- [strings/ToLower](https://golang.org/pkg/strings/#ToLower)
