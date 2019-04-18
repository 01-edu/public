# Compact

## Instructions

Write a function that will take a pointer to a array as parameter and overwrites any element that points to `nil`.

-  If you not sure what the function does. It exists in Ruby. 

## Expected functions

```go
func Compact(ptr *[]string, length int) int {

}
```
## Usage

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
