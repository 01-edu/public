## is-Capitalize

### Instructions 

- Write a function `IsCapitalize` that takes a string and returns `true` if each word in the string begins with an uppercase letter, otherwise, and returns `false`. except for the word that begins with non-alphanumerical characters.
- If the string is empty, return `false`.

### Expected function

```go
func IsCapitalize(s string) bool {

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
	fmt.Println(IsAlpha("Hello! How are you?"))
	fmt.Println(IsAlpha("Hello How Are You"))
	fmt.Println(IsAlpha("Whats 4this 100K?"))
	fmt.Println(IsAlpha("Whatsthis4"))
}
```

And its output :

```console
$ go run .
false
true
true
true
```