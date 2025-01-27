## shoppinglistsort

### Instructions

You were sent to the supermarket with a shopping list. To make your shopping faster, write a function `ShoppingListSort()` that takes a slice of strings and sorts it, according to the string length, returning a slice in which the strings appear in ascending order.

- Strings within the input slice must be of different lengths.

### Expected function

```go
func ShoppingListSort(slice []string) []string {

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
	slice := []string{"Pineapple", "Honey", "Mushroom", "Tea", "Pepper", "Milk"}
	fmt.Println(piscine.ShoppingListSort(slice))
}
```

And its output:

```console
$ go run . | cat -e
[Tea Milk Honey Pepper Mushroom Pineapple]$
```
