## shoppingsummarycounter

### Instructions

You have a receipt from the grocery store and you want to know how many of each item you bought. Write a function that returns this summary.

Given a `string` count the total amount of times each item appears in it and return a summary including the item and its number of appearances as an `int`.

### Expected function

```go
func ShoppingSummaryCounter(str string) map[string]int {

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
	summary := "Burger Water Carrot Coffee Water Water Chips Carrot Carrot Burger Carrot Water"
	for index, element := range piscine.ShoppingSummaryCounter(summary) {
		fmt.Println(index, "=>", element)
	}
}

```

And its output:

```console
$ go run . | cat -e
Burger => 2$
Water => 4$
Carrot => 4$
Coffee => 1$
Chips => 1$
```
