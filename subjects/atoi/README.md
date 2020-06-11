## atoi

### Instructions

- Write a function that simulates the behaviour of the `Atoi` function in Go. `Atoi` transforms a number represented as a `string` in a number represented as an `int`.

- `Atoi` returns `0` if the `string` is not considered as a valid number. For this exercise **non-valid `string` chains will be tested**. Some will contain non-digits characters.

- For this exercise the handling of the signs + or - **does have** to be taken into account.

- This function will **only** have to return the `int`. For this exercise the `error` result of atoi is not required.

### Expected function

```go
func Atoi(s string) int {

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
	fmt.Println(piscine.Atoi("12345"))
	fmt.Println(piscine.Atoi("0000000012345"))
	fmt.Println(piscine.Atoi("012 345"))
	fmt.Println(piscine.Atoi("Hello World!"))
	fmt.Println(piscine.Atoi("+1234"))
	fmt.Println(piscine.Atoi("-1234"))
	fmt.Println(piscine.Atoi("++1234"))
	fmt.Println(piscine.Atoi("--1234"))
}
```

And its output :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
12345
12345
0
0
1234
-1234
0
0
student@ubuntu:~/[[ROOT]]/test$
```
