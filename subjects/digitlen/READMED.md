## digit-len

### Instructions 

Write a function that takes two integers and returns the number of digits in the first integer by the base of the second integer.
- The second integer must be between 2 and 36. if not, the function returns `-1`.
- If the first integer is negative , reverse the sign and count the digits.


### Expected function

```go
func DigitLen(n, base int) int {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import "fmt"

func main(){
    fmt.Println(DigitLen(100, 10))
    fmt.Println(DigitLen(100, 2))
    fmt.Println(DigitLen(-100, 16))
    fmt.Println(DigitLen(100, -1))

}
```

And its output :

```console
$ go run . | cat -e
3
7
2
-1
```
