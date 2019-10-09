## foreach

### Instructions

Write a function `ForEach` that, for an `int` array, applies a function on each elements of that array.

### Expected function

```go
func ForEach(f func(int), arr []int) {

}
```

### Usage

Here is a possible [program](TODO-LINK) to test your function :

```go
package main

import piscine ".."

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	piscine.ForEach(piscine.PrintNbr, arr)
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
123456
student@ubuntu:~/piscine-go/test$
```
