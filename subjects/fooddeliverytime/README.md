## fooddeliverytime

### Instructions

Given the following menu with the corresponding times that each item takes to cook (burger takes 15 min, chips takes 10 min and nuggets takes 12 min), return the time that each order item takes to be prepared and the total amount of time for the order to be ready assuming the items are cooked one after the other.

Write a function `FoodDeliveryTime()` that takes a `string` and returns an `int`.

- Use structs to answer the subject.
- If any of the order items don't exist in the menu above return the error message `404`.

### Expected function

```go
type food struct {
  preptime int
}

func FoodDeliveryTime(order string) int {

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
	fmt.Println(piscine.FoodDeliveryTime("burger"))
	fmt.Println(piscine.FoodDeliveryTime("chips"))
	fmt.Println(piscine.FoodDeliveryTime("nuggets"))
	fmt.Println(piscine.FoodDeliveryTime("burger") + piscine.FoodDeliveryTime("chips") + piscine.FoodDeliveryTime("nuggets"))
}
```

And its output:

```console
$ go run . | cat -e
15$
10$
12$
37$
```
