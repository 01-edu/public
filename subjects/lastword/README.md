## lastword

### Instructions

Write a function `LastWord` that takes a `string` and returns its last word followed by a `\n`.

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
	"fmt"
	"piscine"
)

func main() {
	fmt.Print(piscine.LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(piscine.LastWord(" lorem,ipsum "))
	fmt.Print(piscine.LastWord(" "))
}
```

And its output :

```console
$ go run . | cat -e
not$
lorem,ipsum$
$
$
```

### Notions

- [01-edu/z01](https://github.com/01-edu/z01)
