## foreachprog

### Instructions

Write a function `ForEach` that, for an `int` slice, applies a function on each elements of that slice.

### Expected function

```go
func ForEach(f func(int), a []int) {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	ForEach(PrintNbr, a)
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
