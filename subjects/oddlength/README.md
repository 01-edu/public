## oddlength

### Instructions

Write a function that receives a slice of strings and returns a new slice with the strings in which the length is odd. When the slice is empty or there are no odd strings return an empty slice.

### Expected function

```go
func Oddlength(strings []string) []string {

}
```

### Usage

Here is a possible program to test your function:

```go
package main

import "fmt"

func main() {
    fmt.Println(Oddlength([]string{"Hi", "Hola", "Ola", "Ciao", "Salut", "Hallo"}))
}
```

And its output:

```console
$ go run .
[Ola Salut Hallo]
```
