## makerange

### Instructions

Write a function that takes an `int` min and an `int` max as parameters.
That function returns a slice of `int` with all the values between min and max.

Min is included, and max is excluded.

If min is superior or equal to max, a `nil` slice is returned.

`append` is not allowed for this exercise.

### Expected function

```go
func MakeRange(min, max int) []int {

}
```

### Usage

Here is a possible [program](TODO-LINK) to test your function :

```go
package main

import (
	"fmt"
	piscine ".."
)

func main() {
	fmt.Println(piscine.MakeRange(5, 10))
	fmt.Println(piscine.MakeRange(10, 5))
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
[5 6 7 8 9]
[]
student@ubuntu:~/piscine-go/test$
```
