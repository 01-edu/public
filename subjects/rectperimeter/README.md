## rectperimeter

### Instructions

Write a function that takes two `int`'s as arguments, representing the length of width and height of a rectangle and returning the perimeter of the rectangle.

- If one of the arguments is negative it should return `-1`.

### Expected function

```go
func RectPerimeter(w, h int) int {

}
```

### Usage

Here is a possible program to test your function:

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println(RectPerimeter(10, 2))
	fmt.Println(RectPerimeter(434343, 898989))
	fmt.Println(RectPerimeter(10, -2))
}
```

And its output:

```console
$ go run .
24$
2666664$
-1$
$
```
