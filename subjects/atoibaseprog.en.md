## atoibase

### Instructions

Write a function that takes a `string` number and its `string` base in parameters and returns its conversion as an `int`.

If the base or the `string` number is not valid it returns `0`:

Validity rules for a base :

-   A base must contain at least 2 characters.
-   Each character of a base must be unique.
-   A base should not contain `+` or `-` characters.

Only valid `string` numbers will be tested.

The function **does not have** to manage negative numbers.

### Expected function

```go
func AtoiBase(s string, base string) int {

}
```

### Usage

Here is a possible [program](TODO-LINK) to test your function :

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println(AtoiBase("125", "0123456789"))
	fmt.Println(AtoiBase("1111101", "01"))
	fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(AtoiBase("uoi", "choumi"))
	fmt.Println(AtoiBase("bbbbbab", "-ab"))
}
```

And its output :

```console
student@ubuntu:~/test$ go build
student@ubuntu:~/test$ ./test
125
125
125
125
0
student@ubuntu:~/test$
```
