## islower

### Instructions

Write a function that returns `true` if the `string` passed as the parameter contains only lowercase characters, otherwise, returns `false`.

### Expected function

```go
func IsLower(s string) bool {

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
	fmt.Println(piscine.IsLower("hello"))
	fmt.Println(piscine.IsLower("hello!"))

}
```

And its output :

```console
$ go run .
true
false
$
```
