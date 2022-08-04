## CountStars

### Instructions

Write a function called `CountStars` that counts the stars up to a number given as an argument.

- If the number is negative or equal to 0, return "`No stars`"
- No need to manage overflow.

### Expected Function

```go
func CountStars(num int) string {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import "fmt"

func main() {
	fmt.Println(CountStars(5))
	fmt.Println(CountStars(4))
	fmt.Println(CountStars(-1))
	fmt.Println(CountStars(1))
}
```
And its output :

```console
$ go run . 
1 star...2 stars...3 stars...4 stars...5 stars
1 star...2 stars...3 stars...4 stars
No stars
1 star
```
