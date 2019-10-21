## printnbrinorder

### Instructions

Write a function that prints an `int` passed in parameter.
All possible values of type `int` have to go through, beside negative numbers.
You cannot convert to `int64`.

### Expected function

```go
func PrintNbrInOrder(n int) {

}
```

### Usage

Here is a possible [program](TODO-LINK) to test your function :

```go
package main

import piscine ".."

func main() {
	piscine.PrintNbrInOrder(321)
	piscine.PrintNbrInOrder(0)
	piscine.PrintNbrInOrder(321)
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
1230123
student@ubuntu:~/piscine-go/test$
```
