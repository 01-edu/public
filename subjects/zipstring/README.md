## zipstring

### Instructions

Write a function that takes a `string` and returns a new string that replaces every character with the number of duplicates and the character itself, deleting the extra duplications.

The letters are from the Latin alphabet list only, any other characters, symbols or empty spaces shall not be tested.

### Expected function

```go
func ZipString(s string) string {

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
	fmt.Println(piscine.ZipString("YouuungFellllas"))
}

```

And its output :

```console
$ go run .
1Y1o3u1n1g1F1e4l1a1s
$
```
