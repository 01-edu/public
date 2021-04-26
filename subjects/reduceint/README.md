## reduceint

### Instructions

The function should have as parameters a slice of integers, `a []int` and a function `f func(int, int) int`. You should apply for each element of the slice the arithmetic function, saving it and printing.

### Expected function

```go
func ReduceInt(a []int, f func(int, int) int) {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

func main() {
	mul := func(acc int, cur int) int {
		return acc * cur
	}
	sum := func(acc int, cur int) int {
		return acc + cur
	}
	div := func(acc int, cur int) int {
		return acc / cur
	}
	as := []int{500, 2}
	ReduceInt(as, mul)
	ReduceInt(as, sum)
	ReduceInt(as, div)
}

```

And its output :

```console
student@ubuntu:~/reduceint/test$ go build
student@ubuntu:~/reduceint/test$ ./test
1000
502
250
student@ubuntu:~/reduceint/test$
```
