## setspace

### Instructions

Write a function that takes a PascalCase `string` and returns the same `string` with a space between each word.

- The function must return an empty `string` if the input is an empty `string`.

- The function must return `"Error"` if the input `string` is not a PascalCase `string`.

- The PascalCase begins with a capital letter, and each word begins with a capital letter without a space between them, for example: "HelloWorld" is a valid PascalCase `string`.

- The PascalCase cannot contain any non-alphabetic character, for example: "Hello World12" is not a valid PascalCase `string`.

### Expected function

```go
func SetSpace(s string) string {

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
	fmt.Println(SetSpace("HelloWorld"))
	fmt.Println(SetSpace("HelloWorld12"))
	fmt.Println(SetSpace("Hello World"))
	fmt.Println(SetSpace(""))
	fmt.Println(SetSpace("LoremIpsumWord"))
}

```

And its output:

```console
$ go run . | cat -e
Hello World$
Error$
Error$
$
Lorem Ipsum Word$
$
```
