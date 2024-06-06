## repeatalpha

### Instructions

Write a function called `RepeatAlpha` that takes a `string` and displays it repeating each alphabetical character as many times as its alphabetical index.

`'a'` becomes `'a'`, `'b'` becomes `'bb'`, `'e'` becomes `'eeeee'`, etc...

### Expected Function

```go
func RepeatAlpha(s string) string {
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
	fmt.Println(piscine.RepeatAlpha("abc"))
	fmt.Println(piscine.RepeatAlpha("Choumi."))
	fmt.Println(piscine.RepeatAlpha(""))
	fmt.Println(piscine.RepeatAlpha("abacadaba 01!"))
}
```

And its output:

```console
$ go run . | cat -e
abbccc$
CCChhhhhhhhooooooooooooooouuuuuuuuuuuuuuuuuuuuummmmmmmmmmmmmiiiiiiiii.$
$
abbacccaddddabba 01!$
$
```
