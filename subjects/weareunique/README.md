## Weareunique

### Instructions

Write a function that takes two strings and returns the amount of characters that are not included in both without doubles.

- If there is no unique charachters return -1.
- If both strings are empty  or one of them return -1.


### Expected function

```go

func Weareunique(str1 , str2 string) int {

}

```

### Usage

Here is a possible program to test your function :

```go
package main

import (
	"fmt"
)

func main() {

	arr := [][]string{
		{"foo", "boo"},
		{"abc", "def"},
		{"hello", "yoall"},
		{"everyone", ""},
		{"hello world", "fam"},
	}
	for _, v := range arr {
		fmt.Println(Weareunique(v[0], v[1]))
	}
}
```

And its output :

```console
$ go run .
2
6
4
-1
11
$
```