## activebits

### Instructions

Write a function, ActiveBitsthat, that returns the number of active bits (bits with the value 1) in the binary representation of an integer number.

The function must have the next signature.

### Expected function

```go
func ActiveBits(n int) uint {	

}
```

### Usage

Here is a possible [program](TODO-LINK) to test your function :

```go
package main

import (
  "fmt"
  student ".."
)

func main() {
	nbits := student.ActiveBits(7)
	fmt.Println(nbits)
}
```

And its output :

```console
student@ubuntu:~/student/activebits$ go build
student@ubuntu:~/student/activebits$ ./activebits
10
student@ubuntu:~/student/activebits$ 
```
