## issamestring

### Instructions

Write a function that returns `true` if the two `strings` passed as parameter are equal. Otherwise it returns `false`.

The function will ignore the case of the characters.

### Expected function

```go
func IsSameString(s1, s2 string) bool {

}
```

### Usage

Here is a possible program to test your function:

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println(IsSameString("hello world!", "HELLO WORLD!"))
	fmt.Println(IsSameString("0101 is awesome", "0101 is AWESOME"))
	fmt.Println(IsSameString("foo baz", "foo bar"))
}
```

And its output:

```console
$  go run .
true
true
false
```
