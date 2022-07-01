## sum

### Instructions

Write a function that takes two single digit numbers as a string and returns the sum as an int. The strings can only contain one digit which will be a positive or negative number from 0 to 9. Otherwise, you must return 0.

### Expected function

```go
func Sum(a,b string) int {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import "fmt"

func main() {
	fmt.Println(Sum("1", "2"))
	fmt.Println(Sum("-4", "0"))
	fmt.Println(Sum("7", "-3"))
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
