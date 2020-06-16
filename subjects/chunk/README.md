## chunk

### Instructions

Write a function called `Chunk` that receives as parameters a slice, `slice []int`, and an number `size int`. The goal of this function is to chunk a slice into many sub slices where each sub slice has the length of `size`.

- If the `size` is `0` it should print `\n`

### Expected function

```go
func Chunk(slice []int, size int) {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}
```

And its output :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
[]

[[0 1 2] [3 4 5] [6 7]]
[[0 1 2 3 4] [5 6 7]]
[[0 1 2 3] [4 5 6 7]]
student@ubuntu:~/[[ROOT]]/test$
```
