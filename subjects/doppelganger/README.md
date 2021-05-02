## doppelganger

### Instructions

You are given 2 strings. Find out if first string contains second string. If it does, return index of the first string where second string occurs last time. If it does not contain, return "-1"

### Expected function

```go
package main

func DoppelGanger(s, substr string) int {

}
```

```go
package main

import (
	"fmt"

	"piscine"
)

func main() {
	var result int

	result = piscine.DoppelGanger("aaaaaaa", "a")
	fmt.Println(result) // 6

	result = piscine.DoppelGanger("qwerty", "t")
	fmt.Println(result) // 4

	result = piscine.DoppelGanger("a", "b")
	fmt.Println(result) // -1
}
```
