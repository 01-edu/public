## saveandmiss

### Instructions

Write a function called `SaveAndMiss()` that takes a `string` and an `int` as an argument. The function should move through the `string` in sets determined by the `int`, saving the first set, omitting the second, saving the third, and so on, in a 'save' and 'miss' fashion until the end of the `string` is reached. Return a `string` containing the saved characters.

> If the `int` is `0` or a negative number return the original `string`.

### Expected function

```go
func SaveAndMiss(arg string, num int) string {

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
	fmt.Println(piscine.SaveAndMiss("123456789", 3))
	fmt.Println(piscine.SaveAndMiss("abcdefghijklmnopqrstuvwyz", 3))
	fmt.Println(piscine.SaveAndMiss("", 3))
	fmt.Println(piscine.SaveAndMiss("hello you all ! ", 0))
	fmt.Println(piscine.SaveAndMiss("what is your name?", 0))
	fmt.Println(piscine.SaveAndMiss("go Exercise Save and Miss", -5))
}
```

And its output:

```console
$ go run . | cat -e
123789$
abcghimnostuz$
$
hello you all ! $
what is your name?$
go Exercise Save and Miss$
```
