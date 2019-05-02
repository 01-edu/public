## pointone

### Instructions

- Write a function that takes a **pointer to an int** as argument and gives to this int the value of 1.

### Expected function

```go
func PointOne(n *int) {

}
```

### Usage

Here is a possible [program](TODO-LINK) to test your function :

```go
package main

import (
    "fmt"
    piscine ".."
)

func main() {
    n := 0
    piscine.PointOne(&n)
    fmt.Println(n)
}
```

And its output :

```console
student@ubuntu:~/piscine/test$ go build
student@ubuntu:~/piscine/test$ ./test
1
student@ubuntu:~/piscine/test$
```
