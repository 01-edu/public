## isnegative

### Instructions

Write a [function](TODO-LINK) that prints `'T'` (true) on a single line if the `int` passed as parameter is negative, otherwise it prints `'F'` (false).

### Expected function

```go
func IsNegative(nb int) {

}
```

### Usage

Here is a possible [program](TODO-LINK) to test your function :

```go
package main

import piscine ".."

func main() {
	piscine.IsNegative(1)
	piscine.IsNegative(0)
	piscine.IsNegative(-1)
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
F
F
T
student@ubuntu:~/piscine-go/test$
```
