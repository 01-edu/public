## stringtointarray

### Instructions

Write a function `StringToIntArray()` that takes a string and returns the corresponding integer array.

### Expected function

```go
func StringToIntArray(str string) []int {

}
```

### Usage

Here is a possible program to test your function:

```go
package main

import (
  "fmt"
  "piscine"
)

func main() {
	fmt.Println(piscine.StringToIntArray("A quick brown fox jumps over the lazy dog"))
	fmt.Println(piscine.StringToIntArray("Converted this string into an int"))
	fmt.Println(piscine.StringToIntArray("hello THERE"))
}
```

And its output:

```go
$ go run . | cat -e
[65 32 113 117 105 99 107 32 98 114 111 119 110 32 102 111 120 32 106 117 109 112 115 32 111 118 101 114 32 116 104 101 32 108 97 122 121 32 100 111 103]$
[67 111 110 118 101 114 116 101 100 32 116 104 105 115 32 115 116 114 105 110 103 32 105 110 116 111 32 97 110 32 105 110 116]$
[104 101 108 108 111 32 84 72 69 82 69]$
```
