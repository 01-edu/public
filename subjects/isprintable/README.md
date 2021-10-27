## isprintable

### Instructions

Write a function that returns `true` if the `string` passed as a parameter contains only printable characters, otherwise, returns `false`.

### Expected function

```go
func IsPrintable(s string) bool {

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
	fmt.Println(piscine.IsPrintable("Hello"))
	fmt.Println(piscine.IsPrintable("Hello\n"))

}
```

And its output :

```console
$ go run .
true
false
$
```
