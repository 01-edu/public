## lastword

### Instructions

Write a function `LastWord that takes a `string`and returns its last word followed by a`\n`.

- A word is a section of `string` delimited by spaces or by the start/end of the `string`.

### Expected function

```go
func LastWord(s string) string{

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import (
	"piscine"

	"github.com/01-edu/z01"
)

func main() {
	z01.PrintRune(piscine.LastWord("this        ...       is sparta, then again, maybe    not"))
	z01.PrintRune(piscine.LastWord(" "))
	z01.PrintRune(piscine.LastWord(" lorem,ipsum "))
}
```

And its output :

```console
$ go run . | cat -e
not$
$
lorem,ipsum$
$
```

### Notions

- [01-edu/z01](https://github.com/01-edu/z01)
