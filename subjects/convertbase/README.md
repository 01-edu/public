## convertbase

### Instructions

Write a function that receives three arguments:

- `nbr`: A string representing a numberic value in a [base](https://simple.wikipedia.org/wiki/Base_(mathematics)).

- `baseFrom`: A string representing the base `nbr` it's using.

- `baseTo`: A string representing the base `nbr` should be represented in the returned value.

Only valid bases will be tested.

Negative numbers will not be tested.

### Expected function

```go
func ConvertBase(nbr, baseFrom, baseTo string) string {

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
	result := piscine.ConvertBase("101011", "01", "0123456789")
	fmt.Println(result)
}
```

And its output :

```console
$ go run .
43
$
```
