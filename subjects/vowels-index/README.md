# vowels-index

### Instructions

- Write a function that takes one argument of type string and returns an array of intergers containing the index of the vowels in the string.
- vowels : [a,i ,e ,u o]

### Expected function

```go 
func VowelIdx(str string) []int {
}
```
### Usage

```go
package main

import "fmt"

func main() {

	res := VowelIdx("hello Iyan")
	for _, i := range res {
		fmt.Println(i)
	}

}
```
And its output :
```console

$ go run . 
1 $
4 $
6 $
8 $
```
    
