## alpha-position

## instructions

Write a function named `AlphaPosition` that takes an alphabetical character as a parameter and returns the position of the letter in the alphabet.
- If the character is not in the alphabet, return -1
- If the character is in the alphabet, return the position of the letter in the alphabet
  
### Expected function

```go
func AlphaPosition(c rune) int {
    // your code goes here
}
```

### usage

```go 
package main

import "fmt"

func main(){
    fmt.Println(AlphaPosition('a'))
    fmt.Println(AlphaPosition('z'))
    fmt.Println(AlphaPosition('B'))
    fmt.Println(AlphaPosition('Z'))
    fmt.Println(AlphaPosition('0'))
    fmt.Println(AlphaPosition(' '))
}
```

### output

```console
$ go run . | cat -e
1$
26$
2$
26$
-1$
-1$
```
