## countcharacter

### Instructions

Write a function that, given a `string` and a character as arguments, returns an `int` representing the number of times the character appears in the string.

- If the character is not in the string return 0
- If the string is empty return 0

### Expected Function

```go
func CountChar(str string, c rune) int {

}
```

### Usage

Here is a possible program to test your function:

```go
package main

import "fmt"

func main() {
	fmt.Println(piscine.CountChar("Hello World", 'l'))
	fmt.Println(piscine.CountChar("5  balloons", 5))
	fmt.Println(piscine.CountChar("     ", ' ')) // The first argument contains space and tabulation.
	fmt.Println(piscine.CountChar("The 7 deadly sins", '7'))
}
```

And its output:

```console
$ go run .
3
0
1
1
```