## multorsum

### Instructions

Write a function that receives a slice of int and and int representing the initial value. If the int is odd it should multiply the accumulated value. If it is even, it should be added instead. Return the result. If the slice is empty return 0.

### Expected function

```go
func MultOrSum(ints []int, init int) {
}
```

### Usage

Here is a possible program to test your function :

```go
package main

func main() {
	fmt.Println(MultOrSum([]int{1, 2, 3, 4}, 3 ))
	fmt.Println(MultOrSum([]int{1, 2, 3, 4}, 0 ))
	fmt.Println(MultOrSum([]int{1, -2, 3, 4}, 0 ))
}
```

And its output :

```console
$ go run . | cat -e
19$
10$
-2$
```
