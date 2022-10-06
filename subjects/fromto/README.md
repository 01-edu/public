## fromto

### Instructions

- Write a function that takes two `integers` and returns a `string` representing the range of numbers from the first to the second.
- The function can return the numbers in any order.
- The number must be separated by a comma and a space.
<<<<<<< HEAD

- If any of the arguments is bigger than 99 or less than 0, the function returns `"invalid"` followed by a newline (`'\n'`).

- Add `0` at the first of any number if it is less than 10.

=======
- If any of the arguments is bigger than 99 or less than 0, the function returns `"invalid"` followed by a newline (`'\n'`).
<<<<<<< HEAD
- Add `0` at the first of any number if it is less than 10.  
>>>>>>> docs(fromto):delete space
=======
- Add `0` at the first of any number if it is less than 10.
>>>>>>> style(fromto): white-space, formmating
- Add a new line (`'\n'`) at the end of the string.

### Expected function

```go
func FromTo(from int, to int) string {

}
```

### Usage

Here is a possible program to test your function:

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Print(piscine.FromTo(1, 10))
	fmt.Print(piscine.FromTo(10, 1))
	fmt.Print(piscine.FromTo(10, 10))
	fmt.Print(piscine.FromTo(100, 10))
}
```
And its output:

```console
$ go run . | cat -e
01, 02, 03, 04, 05, 06, 07, 08, 09, 10$
10, 09, 08, 07, 06, 05, 04, 03, 02, 01$
10$
Invalid$
```
