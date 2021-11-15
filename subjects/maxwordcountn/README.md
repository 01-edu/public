## maxwordcountn

### Instructions

Write a function `MaxWordCountN` that will return a `map` of the `n` words that occurs the most in a string `text`. A word is defined as separated by spaces. The `map` you should return will have the word as key and the number of occurences of this word as value.
If two words have the same number of occurences, the one with the lowest ASCII value should be prioritized.

### Expected function

```go
func MaxWordCountN(text string, n int) map[string]int {

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
	fmt.Println(piscine.MaxWordCountN(`
	Orange is the sun sliding to the horizon after a summer day. Orange is the sound of dribbling basketball. Orange is the smell of a tiger lily petal. Orange is the taste of thirst-quenching Nehi Soda. Orange is the color of peach marmalade on a side of toast. Orange is the sound of a carrot popping out of the ground.
	`, 3))
}
```

And its output :

```console
$ go run .
map[Orange:6 of:7 the:8]
$
```
