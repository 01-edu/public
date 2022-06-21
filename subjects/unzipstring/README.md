## Unzipstring


### Instructions

write a function named `Unzipstring` that takes a string in form of a number and an alphabet such us (2a) and returns the original string (aa).

- the number before the alphabet should be between 0 to 9 
- The one  alphabet after each number should be  between a to z or A to Z
- it can be two numbers or two alphabets in a row
- if the Input string doesn't respect the format return `Invalid Input`


### Expected Function

```go
func Unzipstring(s string) string {
    //code here
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
    fmt.Println(piscine.Unzipstring("6H8a"))
    fmt.Println(piscine.Unzipstring("3p6i1W"))
    fmt.Println(piscine.Unzipstring("2O5u2H0e"))
    fmt.Println(piscine.Unzipstring("2f"))
    fmt.Println(piscine.Unzipstring("2a 6p8f"))
    fmt.Println(piscine.Unzipstring("2p4;7g"))
    fmt.Println(piscine.Unzipstring("2t4dD"))
    fmt.Println(piscine.Unzipstring("82t4D"))
    fmt.Println(piscine.Unzipstring(""))
}
```

And its output :

```console
$ go run .
$HHHHHaaaaaaaa
$pppiiiiiiWWW
$OOuuuuuHH
$ff
$Invalid Input
$Invalid Input
$Invalid Input
$Invalid Input
$Invalid Input
```
