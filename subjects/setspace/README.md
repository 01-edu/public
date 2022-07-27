## set-space

### Instructions 

Write a function that takes a pascal case string and returns the same string with spaces between each word.

- The function must return an empty string if the input string is empty.

- The function must return `"Error"` if the input string is not a pascal case string.

- The pascal case begins with a capital letter, and each word begins with a capital letter without a space between them for example: `"HelloWorld`"` is a valid pascal case string.

- The pascal case cannot contain any non-alphabetic character, for example: `"Hello World12`"` is not a valid pascal case string.

### Expected function

```go
func SetSpace(s string) string {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import "fmt"

func main(){
    fmt.Println(SetSpace("HelloWorld"))
    fmt.Println(SetSpace("Hello World12"))
    fmt.Println(SetSpace("Hello World"))
    fmt.Println(SetSpace("LoremIpSumWord"))
    fmt.Println(SetSpace(""))
}
```

And its output :

```console
$ go run . | cat -e
Hello World$
Error$
Error$
Lorem Ipsum Word$
$
```
