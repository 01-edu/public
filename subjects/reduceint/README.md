## reduceint

### Instructions

The function should have as parameters a slice of integers `a []int` and a function `f func(int, int) int`.  For each element of the slice, it should apply the function `f func(int, int) int`, save the result and then print it.

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
$ go run .
1000
502
250
$
```
