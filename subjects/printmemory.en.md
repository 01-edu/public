## printmemory

### Instructions

Write a function that takes `(arr [10]int)`, and displays the memory as in the example.

### Expected function

```go
func PrintMemory(arr [10]int) {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

func main() {
	PrintMemory([10]int{104, 101, 108, 108, 111, 16, 21, 42})
}
```

And its output :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test | cat -e
6800 0000 6500 0000 6c00 0000 6c00 0000 $
6f00 0000 1000 0000 1500 0000 2a00 0000 $
0000 0000 0000 0000 $
hello..*..$
student@ubuntu:~/[[ROOT]]/test$
```
