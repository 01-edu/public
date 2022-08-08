# printif

### Instructions

Create a function  `PrintIf()` that takes a string in the parameter and prints`G` followed bye newline if the argument length is more than 3; otherwise, print `\n`.

- If it's an empty string returns "Invalid Output\n"

### Expected function

```go
func PrintIf(str string) string {
}
```

### Usage

Here is a possible program to test your function:

```go
package main

import (
	"fmt"
)

func main() {
    fmt.Print(PrintIf("abcdefz"))
    fmt.Print(PrintIf("T2d97d1"))
    fmt.Print(PrintIf("14"))
    fmt.Print(PrintIf("14"))
}
```

And its output :

```go
$ go run . | cat -e
G$
G$
$
Invalid Output$
```
