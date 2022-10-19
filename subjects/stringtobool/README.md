## stringtobool

### Instructions

Write a function that takes a `string` as an argument and returns a `boolean`.

- If the `string` equals `True` or `T` or `t` return `true` otherwise, return `false`.


### Expected function

```go
func StringToBool(s string) bool {

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
	fmt.Println(piscine.StringToBool("True"))
	fmt.Println(piscine.StringToBool("T"))
	fmt.Println(piscine.StringToBool("False"))
	fmt.Println(piscine.StringToBool("TTFF"))
}
```

And its output:

```console
$ go run .
true
true
false
false
```
