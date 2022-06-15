## vowels-index

### Instructions

Write a function that takes one argument of type string and returns an array of integers containing the index of the vowels in the string.
- vowels: a, e, i, o, u]

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

	res := VowelsIndex("hello Iyan")
	for _, i := range res {
		fmt.Println(i)
	}

}

```
And its output :
```console
$ go run .  | cat -e
1$
4$
6$
8$
```
    
