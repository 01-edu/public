## basicatoi

### Instructions

- Write a function that simulates the behaviour of the `Atoi` function in Go. `Atoi` transforms a number defined as a `string` in a number defined as an `int`.

- Atoi returns `0` if the `string` is not considered as a valid number. For this exercise **only valid** `string` will be tested. They will only contain one or several digits as characters.

- For this exercise the handling of the signs `+` or `-` does not have to be taken into account.

- This function will **only** have to return the `int`. For this exercise the `error` return of `Atoi` is not required.

### Expected function

```go
func BasicAtoi(s string) int {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import (
	"fmt"
	piscine ".."
)

func main() {
	fmt.Println(piscine.BasicAtoi("12345"))
	fmt.Println(piscine.BasicAtoi("0000000012345"))
	fmt.Println(piscine.BasicAtoi("000000"))
}
```

And its output :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
12345
12345
0
student@ubuntu:~/[[ROOT]]/test$
```
