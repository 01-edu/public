## atoi

### Instructions

-   Write a [function](TODO-LINK) that simulates the behaviour of the `Atoi` function in Go. `Atoi` transforms a number represented as a `string` in a number represented as an `int`.

-   `Atoi` returns `0` if the `string` is not considered as a valid number. For this exercise **non-valid `string` chains will be tested**. Some will contain non-digits characters.

-   For this exercise the handling of the signs + or - **does have** to be taken into account.

-   This function will **only** have to return the `int`. For this exercise the `error` result of atoi is not required.

### Expected function

```go
func Atoi(s string) int {

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
	s2 := "0000000012345"
	s3 := "012 345"
	s4 := "Hello World!"
	s5 := "+1234"
	s6 := "-1234"
	s7 := "++1234"
	s8 := "--1234"

	n := piscine.Atoi(s)
	n2 := piscine.Atoi(s2)
	n3 := piscine.Atoi(s3)
	n4 := piscine.Atoi(s4)
	n5 := piscine.Atoi(s5)
	n6 := piscine.Atoi(s6)
	n7 := piscine.Atoi(s7)
	n8 := piscine.Atoi(s8)

	fmt.Println(n)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)
	fmt.Println(n5)
	fmt.Println(n6)
	fmt.Println(n7)
	fmt.Println(n8)
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
12345
12345
0
0
1234
-1234
0
0
student@ubuntu:~/piscine-go/test$
```
