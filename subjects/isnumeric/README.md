## isnumeric

### Instructions

Write a function that returns `true` if the `string` passed as a parameter contains only numerical characters, otherwise, returns `false`.

### Expected function

```go
func IsNumeric(s string) bool {

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
	fmt.Println(piscine.IsNumeric("010203"))
	fmt.Println(piscine.IsNumeric("01,02,03"))
}
```

And its output :

```console
$ go run .
true
false
$
```
