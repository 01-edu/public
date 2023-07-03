## stringtobool

### Instructions

Write a function that takes a `string` as an argument and returns a `boolean`.

- If the `string` equals `True`, `T` or `t` return `true`, otherwise return `false`.

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
)

func main() {
	fmt.Println(StringToBool("True"))
	fmt.Println(StringToBool("T"))
	fmt.Println(StringToBool("False"))
	fmt.Println(StringToBool("TTFF"))
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
