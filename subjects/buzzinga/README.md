# buzzinga

### Instructions
Write a function named `BuzZinga()` that takes a number as an argument and prints the following:
    
    If the number is divisible by 3, print Buz followed by newline.
    If the number is divisible by 5, print Zinga followed by newline.
    If the number is divisible by 3 and 5, print BuzZinga followed by newline.
    If the number is not divisible by 3 or 5, print * followed by newline.

### Expected function

```go
func BuzZinga(number int) {

}
```
### Usage
```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.BuzZinga(15))
}
```
And its output :

```go
$ go run . | cat -e
*$
*$
Buz$
*$
Zinga$
Buz$
*$
*$
Buz$
Zinga$
*$
Buz$
*$
*$
BuzZinga$
```