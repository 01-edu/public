## printmemory

### Instructions

Write a function that takes `(arr [10]int, size int)`, and displays the memory as in the example.

Your function must be declared as follows :

```go
func PrintMemory(arr [10]int, size int) {
	
}
```

### Usege

Here is a possible program to test your function and it's output :

```go
func main() {
	arr := [10]int{0, 23, 150, 255, 12, 16, 21, 42}
	PrintMemory(arr, len(arr))
}
```

```console
student@ubuntu:~/piscine/test$ go build
student@ubuntu:~/piscine/test$ ./test
0000 0000 1700 0000 9600 0000 ff00 0000 ................
0c00 0000 1000 0000 1500 0000 2a00 0000 ............*...
0000 0000 0000 0000 					........
student@ubuntu:~/piscine/test$
```