## isnegative

Write a function named `IsNegative()` that defines if a number (You should check if it's a number) is negative or positive.
- Your function print `P` if the number is positive
-  Your function print `F` if the number is negative
-  For the number is zero print `0`.
-  If it's not a number print newline.
-  Your program should always print `('\n') `at the end of the output.

### Expected function
```go
func StrisNegative(str string) {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import "piscine"

func main() {
	piscine.StrisNegative("585")
	piscine.StrisNegative("-58")
	piscine.StrisNegative("55s44")
	piscine.StrisNegative("101-1331")
	piscine.StrisNegative("5544-")
}
```

And its output :

```console
$ go run .
P
N
!
!
!
```

