# fooddeliverytime

### Instructions

Given a menu with the corresponding time that each item takes to cook (burger takes 15 min, chips takes 10 min and nuggets takes 12 min), return the time each item takes to be prepared and the total amount of time for the order to be ready assuming the items are cooked one after the other.

Write a function `FoodDeliveryTime()` that takes a `string` and returns an `int`.

- Use structures to answer the subject.

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

import "fmt"

func main() {
  fmt.Println(FoodDeliveryTime("burger"))
  fmt.Println(FoodDeliveryTime("chips"))
  fmt.Println(FoodDeliveryTime("nuggets"))
  fmt.Println(FoodDeliveryTime("burger") + FoodDeliveryTime("chips") + FoodDeliveryTime("nuggets"))
}
```

And its output:

```go
$ go run . | cat -e
15$
10$
12$
37$
```
