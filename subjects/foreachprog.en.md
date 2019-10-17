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

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	ForEach(PrintNbr, arr)
}

func PrintNbr(int) {

}
```

And its output :

```console
student@ubuntu:~/test$ go build
student@ubuntu:~/test$ ./test
123456
student@ubuntu:~/test$
```
