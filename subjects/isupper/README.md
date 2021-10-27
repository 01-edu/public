## isupper

### Instructions

Write a function that returns `true` if the `string` passed as parameter contains only uppercase characters, otherwise, returns `false`.

### Expected function

```go
func IsUpper(s string) bool {

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
	fmt.Println(piscine.IsUpper("HELLO"))
	fmt.Println(piscine.IsUpper("HELLO!"))

}
```

And its output :

```console
$ go run .
true
false
$
```
