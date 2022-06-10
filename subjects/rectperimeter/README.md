# rectperimeter

### Instructions

write a function that accepts two integer arguments, representing the two sides length of a rectangle and returning the perimeter of the rectangle.

your function should return `-1` if one of the arguments is negative.

### Expected function

```go
func RectPerimeter(l, w int) int {

}
```

### Usage


```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.RectPerimeter(10, 2))
	fmt.Println(piscine.RectPerimeter(434343, 898989))
}
```
