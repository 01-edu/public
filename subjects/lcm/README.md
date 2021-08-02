## lcm

### Instructions

Write a function, `lcm`, that returns least common multiple.

It will be tested with positive `int` values and `0`.

### Expected function

```go
func Lcm(first, second int) int {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import (
	"fmt"
	"student"
)

func main() {
	fmt.Println(student.Lcm(2, 7))
	fmt.Println(student.Lcm(0, 4))
}
```

### Output

```console
$ go run .
14
0
$
```
