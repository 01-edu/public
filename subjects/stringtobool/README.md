## string-to-bool

### Instructions

- Write a function that takes a string in parameter and returns: 

- If the string equals `True` or `T` or `t` returns `true` otherwise, returns `false`.


### Expected function

```go
func StringToBool(s string) string {

}
```

### Usage

Here is a possible program to test your function :

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

And its output :

```console
$ go run .
true
true
false
false
```
