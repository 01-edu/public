## hashcode

### Instructions

- Write a function called HashCode(string) that takes a string in parameter and returns a new string hashed 

- Hash equation: (ASCII of current character + size of the string) % 127, so it can be in the limit of ASCII size '127'.

- If the final number gives an unprintable character, add 33.


### Expected function

```go
func HashCode(dec string) string {
}
```

### Usage

Here is a possible program to test your function :

```go
package main

import "fmt"

func main() {
    fmt.Println(HashCode("A")) 
    fmt.Println(HashCode("AB"))
    fmt.Println(HashCode("BAC"))
    fmt.Println(HashCode("Hello World")) 
}
```

And its output :

```
$ go run . 
B
CD
EDF
Spwwz+bz}wo
```