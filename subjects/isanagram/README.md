## isanagram

### Instructions

Write a function that returns `true` if two strings are anagrams, otherwise returns `false`.
An anagram is a string made by using the letters of another string in a different order.

Only lower case characters will be given.

### Expected function

```go
package piscine

func IsAnagram(str1, str2 string) bool {

}
```

### Usage

Here is a possible program to test your function:

```go
package main

import (
	"piscine"
	"fmt"
)

func main() {
	test1 := piscine.IsAnagram("listen", "silent")
	fmt.Println(test1)

	test2 := piscine.IsAnagram("alem", "school")
	fmt.Println(test2)

	test3 := piscine.IsAnagram("neat", "a net")
	fmt.Println(test3)

	test4 := piscine.IsAnagram("anna madrigal", "a man and a girl")
	fmt.Println(test4)
}
```

Its output:

```console
$ go run .
true
false
true
true
$
```
