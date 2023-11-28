## retainfirsthalf

### Instructions

Write a function called `RetainFirstHalf()` that takes a `string` as an argument and returns the first half of this `string`.

- If the length of the `string` is odd, round it down.
- If the `string` is empty, return an empty string.
- If the `string` length is equal to one, return the string.

### Expected function

```go
func RetainFirstHalf(str string) string {

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
	fmt.Println(piscine.RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
	fmt.Println(piscine.RetainFirstHalf("A"))
	fmt.Println(piscine.RetainFirstHalf(""))
	fmt.Println(piscine.RetainFirstHalf("Hello World"))
}
```

And its output:

```console
$ go run . | cat -e
This is the 1st half$
A$
$
Hello$
```
