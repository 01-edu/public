## iscapitalized

### Instructions

Write a function `IsCapitalized()` that takes a `string` as an argument and returns `true` if each word in the `string` begins with either an uppercase letter or a non-alphabetic character.

- If any of the words begin with a lowercase letter return `false`.
- If the `string` is empty return `false`.

### Expected function

```go
func IsCapitalized(s string) bool {

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
	fmt.Println(IsCapitalized("Hello! How are you?"))
	fmt.Println(IsCapitalized("Hello How Are You"))
	fmt.Println(IsCapitalized("Whats 4this 100K?"))
	fmt.Println(IsCapitalized("Whatsthis4"))
	fmt.Println(IsCapitalized("!!!!Whatsthis4"))
	fmt.Println(IsCapitalized(""))
}
```

And its output:

```console
$ go run .
false
true
true
true
true
false
```
