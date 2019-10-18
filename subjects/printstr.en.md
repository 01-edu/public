## printstr

### Instructions

-   Write a function that prints one by one the characters of a `string` on the screen.

### Expected function

```go
func PrintStr(str string) {

}
```

### Hints

Here is a possible [program](TODO-LINK) to test your function :

```go
package main

import piscine ".."

func main() {
	str := "Hello World!"
	piscine.PrintStr(str)
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test | cat -e
Hello World!%
student@ubuntu:~/piscine-go/test$
```
