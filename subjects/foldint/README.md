## foldint

### Instructions

Write a function called `FoldInt` that simulates the behavior of reduce from JavaScript.

The function should have as parameters a function, `f func(int, int) int` a slice of integers, `slice []int` and an int `acc int`. You should apply for each element of the slice the arithmetic function, saving and printing it. The function will be tested with our own functions `Add, Sub, and Mul`.

### Expected function

```go
func FoldInt(f func(int, int) int, a []int, n int) {

}
```

### Usage

Here is a possible program to test your function:

```go
package main

func main() {
	table := []int{1, 2, 3}
	ac := 93
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
	fmt.Println()

	table = []int{0}
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
}

```

And its output :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
99
558
87

93
0
93
student@ubuntu:~/[[ROOT]]/test$
```
