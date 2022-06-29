## sum

### Instructions

Write a function that takes two single digit numbers as a string and returns the sum as an int. If one of the arguments is not a single digit number, return 0;

### Expected function

```go
func Sum(a,b string) int {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import (
    "piscine"
    "fmt"
)

func main() {
	fmt.Println(piscine.Sum("1", "2"))
	fmt.Println(piscine.Sum("-4", "0"))
	fmt.Println(piscine.Sum("7", "-3"))
}
```

And its output :

```console
$ go run . | cat -e
3$
-4$
4$
$
```
