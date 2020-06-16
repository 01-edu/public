## reduceint

### Instructions

Write a function called `reduceint` that simulates the behaviour of reduce from JavaScript.

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
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
1000
502
250
student@ubuntu:~/[[ROOT]]/test$
```
