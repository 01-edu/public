## byebyefirst

### Instructions

Write a function that takes a slice of `string`'s and returns a new slice without the first element.

- If the slice is empty, return the empty slice.

### Expected function

```go
func ByeByeFirst(strings []string) []string {

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
	fmt.Println(piscine.ByeByeFirst([]string{}))
	fmt.Println(piscine.ByeByeFirst([]string{"one arg"}))
	fmt.Println(piscine.ByeByeFirst([]string{"first", "second"}))
	fmt.Println(piscine.ByeByeFirst([]string{"", "abcd", "efg"}))
}
```

And its output:

```console
$ go run . | cat -e
[]$
[]$
[second]$
[abcd efg]$
```
