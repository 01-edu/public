## trimatoi

### Instructions

- Write a function that transforms a number within a `string` in a number represented as an `int`.

- For this exercise the handling of the signs + or - **has** to be taken into account. If one of the signs is encountered before any number it should determine the sign of the returned `int`.

- This function will **only** have to return the `int`. In case of invalid input, the function should return `0`.

- Note: There will never be more than one sign by string in the tests.

### Expected function

```go
func TrimAtoi(s string) int {

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
	fmt.Println(piscine.TrimAtoi("12345"))
	fmt.Println(piscine.TrimAtoi("str123ing45"))
	fmt.Println(piscine.TrimAtoi("012 345"))
	fmt.Println(piscine.TrimAtoi("Hello World!"))
	fmt.Println(piscine.TrimAtoi("sd+x1fa2W3s4"))
	fmt.Println(piscine.TrimAtoi("sd-x1fa2W3s4"))
	fmt.Println(piscine.TrimAtoi("sdx1-fa2W3s4"))
}
```

And its output :

```console
student@ubuntu:~/trimatoi/test$ go build
student@ubuntu:~/trimatoi/test$ ./test
12345
12345
12345
0
1234
-1234
1234
student@ubuntu:~/trimatoi/test$
```
