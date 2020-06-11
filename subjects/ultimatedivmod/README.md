## ultimatedivmod

### Instructions

- Write a function that will be formatted as below.

### Expected function

```go
func UltimateDivMod(a *int, b *int) {

}
```

- This function will divide the int **a** and **b**.
- The result of this division will be stored in the int pointed by **a**.
- The remainder of this division will be stored in the int pointed by **b**.

### Usage

Here is a possible program to test your function :

```go
package main

import (
	"fmt"
	piscine ".."
)

func main() {
	a := 13
	b := 2
	piscine.UltimateDivMod(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
```

And its output :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
6
1
student@ubuntu:~/[[ROOT]]/test$
```
