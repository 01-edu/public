## split

### Instructions

Write a function that receives a string and a separator and returns a `slice of strings` that results of splitting the string `s` by the separator `sep`.

### Expected function

```go
func Split(s, sep string) []string {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import "fmt"

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", Split(s, "HA"))
}
```

And its output :

```console
$ go run .
[]string{"Hello", "how", "are", "you?"}
$
```
