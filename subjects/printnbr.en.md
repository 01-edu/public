## printnbr

### Instructions

Write a function that prints an `int` passed in parameter.
All possible values of type `int` have to go through.
You cannot convert to `int64`.

### Expected function

```go
func PrintNbr(n int) {

}
```

### Usage

Here is a possible [program](TODO-LINK) to test your function :

```go
package main

import piscine ".."

func main() {
	piscine.PrintNbr(-123)
	piscine.PrintNbr(0)
	piscine.PrintNbr(123)
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
-1230123
student@ubuntu:~/piscine-go/test$
```
