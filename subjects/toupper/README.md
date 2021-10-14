## toupper

### Instructions

Write a function that capitalizes each letter in a `string`.

### Expected function

```go
func ToUpper(s string) string {

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
	fmt.Println(piscine.ToUpper("Hello! How are you?"))
}
```

And its output :

```console
$ go run .
HELLO! HOW ARE YOU?
$
```

### Notions

- [strings/ToUpper](https://golang.org/pkg/strings/#ToUpper)
