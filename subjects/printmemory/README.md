## printmemory

### Instructions

Write a function that takes `(arr [10]byte)`, and displays the memory as in the example.

After displaying the memory the function must display all the ASCII graphic characters. The non printable characters must be replaced by a dot.

The ASCII graphic characters are any characters intended to be written, printed, or otherwise displayed in a form that can be read by humans, present on the ASCII encoding.

### Expected function

```go
func PrintMemory(arr [10]byte) {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import "piscine"

func main() {
	piscine.PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
```

And its output :

```console
$ go run . | cat -e
68 65 6c 6c$
6f 10 15 2a$
00 00$
hello..*..$
$
```
