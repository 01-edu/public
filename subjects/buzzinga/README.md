# buzzinga

### Instructions
Write a function named `BuzZinga()` that takes a number as an argument and prints the following:
    
    If the number is divisible by 3, print Buz followed by a newline.
    If the number is divisible by 5, print Zinga followed by a newline.
    If the number is divisible by 3 and 5, print BuzZinga followed by a newline.
    If the number is not divisible by 3 or 5, print * followed by a newline.

### Expected function

```go
func BuzZinga(number int) {

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
	fmt.Println(BuzZinga(15))
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