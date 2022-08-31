# shoppinglistsort

### Instructions

You were sent to the supermarket with a shopping list. To make your shopping faster, write a function `ShoppingListSort()` that takes an array of strings and sorts it, according to the string length, returning an array in which the strings appear in ascending order.

In the case of equal string length the order does not matter.

### Expected function

```go
func ShoppingListSort(array[] string)[]string {

}
```

### Usage

Here is a possible program to test your function:

```go
package main

import "fmt"

func main() {
	 array:= []string{"Banana", "Mushroom", "Salt", "Pepper","Tea", "Milk"}
   fmt.Println(ShoppingListSort(array))
}
```

And its output:

```go
$ go run . | cat -e
[Tea Salt Milk Banana Pepper Mushroom]$
```
