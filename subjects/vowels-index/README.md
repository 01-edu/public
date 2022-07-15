## vowels-index

### Instructions

Write a function that takes one argument of type string and returns an array of integers containing the index of the vowels in the string.
- vowels: [a, e, i, o, u]

### Expected function

```go 
func VowelsIndex(str string) []int {
}
```
### Usage

Here is a possible program to test your function

```go
package main

import "fmt"

func main() {
	res := []string{"student", "hello Iyan","bcdfgh", "wOrld", "","AAEO$o;jw"}
	for _, i := range res {
		fmt.Println(VowelIdx(i))
	}
}
```
And its output :

```console
$ go run .
[2 4]
[1 4 6 8]
[]
[1]
[]
[0 1 2 3 5]
```
    
