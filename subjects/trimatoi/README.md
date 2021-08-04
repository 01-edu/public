## trimatoi

### Instructions

- Write a function that transforms a number within a `string`, in an `int`.

- For this exercise the handling of the `-` sign **has** to be taken into account. If the sign is encountered before any number it should determine the sign of the returned `int`.

- This function will **only** return an `int`. In case of invalid input, the function should return `0`.

- **Note**: There will never be more than one sign by `string` in the tests.

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
	"piscine"
)

func main() {
	fmt.Println(piscine.TrimAtoi("12345"))
	fmt.Println(piscine.TrimAtoi("str123ing45"))
	fmt.Println(piscine.TrimAtoi("012 345"))
	fmt.Println(piscine.TrimAtoi("Hello World!"))
	fmt.Println(piscine.TrimAtoi("sd+x1fa2W3s4"))
	fmt.Println(piscine.TrimAtoi("sd-x1fa2W3s4"))
	fmt.Println(piscine.TrimAtoi("sdx1-fa2W3s4"))
	fmt.Println(piscine.TrimAtoi("sdx1+fa2W3s4"))
}
```

And its output :

```console
$ go run .
12345
12345
12345
0
1234
-1234
1234
1234
$
```
