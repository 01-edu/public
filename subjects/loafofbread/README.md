# loafofbread

### Instructions

Write a function `LoafOfBread()` that takes a string and returns another one with words of 5 characters and skips the next character followed by newline `\n`.

- If there is a space in the middle of a word it should ignore it and get the first character until getting to a length of 5.
- If the string less than 5 characters returns "Invalid Output\n"
-
### Expected function

```go
func LoafOfBread(str string) string {
}
```

### Usage

Here is a possible program to test your function :

```go
package main

import "fmt"

func main() {
	LoafOfBread("deliciousbread")
  LoafOfBread("This is a loaf of bread")
  LoafOfBread("Bread crumbles")
  fmt.Print(LoafOfBread("deliciousbread"))
  fmt.Print(LoafOfBread("This is a loaf of bread"))
  fmt.Print(LoafOfBread("loaf"))
}
```

And its output :

```go
$ go run . | cat -e
delic ousbr ad$
Thisi ashor sente ce$
This  s a l af of bread$
Invalid Output$
```
