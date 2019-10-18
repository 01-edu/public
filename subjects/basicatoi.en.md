## basicatoi

### Instructions

-   Write a function that simulates the behaviour of the atoi function in Go. Atoi transforms a number defined as a `string` in a number defined as an `int`.

-   Atoi returns `0` if the `string` is not considered as a valid number. For this exercise **only valid** `string` chains will be tested. They will only contain one or several digits as characters.

-   For this exercise the handling of the signs `+` or `-` does not have to be taken into account.

-   This function will **only** have to return the `int`. For this exercise the `error` return of atoi is not required.

### Expected function

```go
func BasicAtoi(s string) int {

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
	s3 := "000000"

	n := piscine.BasicAtoi(s)
	n2 := piscine.BasicAtoi(s2)
	n3 := piscine.BasicAtoi(s3)

	fmt.Println(n)
	fmt.Println(n2)
	fmt.Println(n3)
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
12345
12345
0
student@ubuntu:~/piscine-go/test$
```
