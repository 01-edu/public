## weareunique

### Instructions

Write a function that takes two strings and returns the number of characters that are not included in both, without doubles.
- If there is no unique characters return `0`.
- If both strings are empty return `-1`.

### Expected function

```go
func WeAreUnique(str1 , str2 string) int {

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
	arr := [][]string{
		{"foo", "boo"},
		{"", ""},
		{"abc", "def"},
		{"hello", "yoall"},
		{"everyone", ""},
		{"hello world", "fam"},
		{"pomme", "pomme"},
		{"", "exam"},
	}
	for _, v := range arr {
		fmt.Println(WeAreUnique(v[0], v[1]))
	}
}
```

And its output:

```console
$ go run .
2
-1
6
4
6
11
0
4
```
