## Unzipstring

### Instructions

write a function named `Unzipstring` that takes a string in form of a number and printable character that is not a number and returns the original string.

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
)

func main() {
    fmt.Println(Unzipstring("2f"))
    fmt.Println(Unzipstring("2O5u2H0e"))
    fmt.Println(Unzipstring("3p6i1W"))
    fmt.Println(Unzipstring("6H8a"))
    fmt.Println(Unzipstring("2p4;7g"))
    fmt.Println(Unzipstring("2a 6p8f"))
    fmt.Println(Unzipstring("2t4dD"))
    fmt.Println(Unzipstring("82t4D"))
    fmt.Println(Unzipstring(""))
}
```

And its output :

```console
$ go run .
$ff
$OOuuuuuHH
$pppiiiiiiWWW
$HHHHHaaaaaaaa
$pp;;;;ggggggg
$Invalid Input
$Invalid Input
$Invalid Input
$Invalid Input
```
