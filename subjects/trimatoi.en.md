## trimatoi

### Instructions

-   Write a [function](TODO-LINK) that transforms a number within a `string` in a number represented as an `int`.

-   For this exercise the handling of the signs + or - **has** to be taken into account. If one of the signs is encountered before any number it should determine the sign of the returned `int`.

-   This function will **only** have to return the `int`. In case of invalid input, the function should return `0`.

### Expected function

```go
func TrimAtoi(s string) int {

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
	s := "12345"
	s2 := "str123ing45"
	s3 := "012 345"
	s4 := "Hello World!"
	s5 := "sd+x1fa2W3s4"
        s6 := "sd-x1fa2W3s4"
        s7 := "sdx1-fa2W3s4"

	n := piscine.TrimAtoi(s)
	n2 := piscine.TrimAtoi(s2)
	n3 := piscine.TrimAtoi(s3)
	n4 := piscine.TrimAtoi(s4)
	n5 := piscine.TrimAtoi(s5)
	n6 := piscine.TrimAtoi(s6)
	n7 := piscine.TrimAtoi(s7)

	fmt.Println(n)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)
	fmt.Println(n5)
	fmt.Println(n6)
	fmt.Println(n7)
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
12345
12345
12345
0
1234
-1234
1234
student@ubuntu:~/piscine-go/test$
```
