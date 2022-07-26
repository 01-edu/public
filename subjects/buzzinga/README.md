# buzzinga

### Instructions

Write a function called `BuzZinga()` that takes a number as an argument, then loops over it and does the following for each iteration:
- If the number is divisible by 3, print `Buz` followed by a newline
- If the number is divisible by 5, print `Zinga` followed by a newline
- If the number is divisible by 3 and 5, print `BuzZinga` followed by a newline
- If the number is negative or not divisible by 3 or 5, print `*` followed by a newline.
- if the number is 0, print `BuzZinga` followed by a newline.

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
	BuzZinga(15)
    fmt.Println("----------")
	BuzZinga(-3)
    fmt.Println("----------")
    BuzZinga(5)
    fmt.Println("----------")
    BuzZinga(0)
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
----------
*$
----------
*$
*$
Buz$
*$
Zinga$
----------
BuzZinga$
```
