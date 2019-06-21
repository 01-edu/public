## compact

### Instructions

Write a function `Compact` that takes a pointer to an array as parameter and overwrites the elements that points to `nil`.

- Hint: This fonction exists in Ruby.

### Expected function

```go
func Compact(ptr *[]string, length int) int {

}
```

### Usage

Here is a possible [program](TODO-LINK) to test your function :

```go
package main

import fmt

func main() {
	array := []string{"hello", " ", "there", " ", "bye"}

	ptr := &array
	fmt.Println(Compact(ptr, len(array)))
}
```

And its output :

```console
student@ubuntu:~/piscine/test$ go build
student@ubuntu:~/piscine/test$ ./test
3
student@ubuntu:~/piscine/test$
```
