## divmod

### Instructions

-   Write a function that will be formatted as below.

### Expected function

```go
func DivMod(a int, b int, div *int, mod *int) {

}
```

-   This function will divide the int **a** and **b**.
-   The result of this division will be stored in the int pointed by **div**.
-   The remainder of this division will be stored in the int pointed by **mod**.

### Usage

Here is a possible [program](TODO-LINK) to test your function :

```go
package main

import (
    "fmt"
    piscine ".."
)

func main() {
	a := 13
	b := 2
	var div int
	var mod int
	piscine.DivMod(a, b, &div, &mod)
	fmt.Println(div)
	fmt.Println(mod)
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
6
1
student@ubuntu:~/piscine-go/test$
```
