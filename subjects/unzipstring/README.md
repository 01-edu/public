## Unzipstring

### Instructions

Write a function called `Unzipstring` that takes a string that will be a kind of code, and your function will have to decrypt it and return a new string with the output.
It works as follows:
The string will be formed by a number followed by a letter, and the purpose is to print this letter the number of times that is requested by the number that precedes it.
   
    "3w" ==> www
    "2m3e" ==> mmee

- The letter after each number must be between `a` and `z` or `A` and `Z`.
- The number before the letter must be between `0` and `9`.
- You cannot have two numbers or two letters in a row.
- If the input string does not respect the format, return `Invalid Input`.

### Expected Function

```go
func Unzipstring(s string) string {
}
```

### Usage

Here is a possible program to test your function:

```go
package main

import "fmt"

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
$pppiiiiiiW
$HHHHHHaaaaaaaa
$Invalid Input
$Invalid Input
$Invalid Input
$Invalid Input
$Invalid Input
```
